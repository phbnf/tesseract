package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/transparency-dev/formats/log"
	tdnote "github.com/transparency-dev/formats/note"
	"github.com/transparency-dev/tessera/api/layout"
	"github.com/transparency-dev/tesseract/internal/client"
	"github.com/transparency-dev/tesseract/internal/types/staticct"
	"golang.org/x/mod/sumdb/note"
	"k8s.io/klog/v2"
)

var (
	monitoringURL = flag.String("monitoring_url", "", "Base tlog-tiles URL")
	leafIndex     = flag.String("leaf_index", "", "The index of the leaf to fetch")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	if *monitoringURL == "" {
		klog.Exitf("--monitoring_url must be set")
	}
	if *leafIndex == "" {
		klog.Exitf("--leaf_index must be set")
	}
	li, err := strconv.ParseUint(*leafIndex, 10, 64)
	if err != nil {
		klog.Exitf("Invalid --leaf_index: %v", err)
	}

	logURL, err := url.Parse(*monitoringURL)
	if err != nil {
		klog.Exitf("Invalid --monitoring_url %q: %v", *monitoringURL, err)
	}
	hc := &http.Client{
		Timeout: 30 * time.Second,
	}
	fetcher, err := client.NewHTTPFetcher(logURL, hc)
	if err != nil {
		klog.Exitf("Failed to create HTTP fetcher: %v", err)
	}

	ctx := context.Background()

	cpRaw, err := fetcher.Checkpoint(ctx)
	if err != nil {
		klog.Exitf("Failed to fetch checkpoint: %v", err)
	}

	// We need a verifier to parse the checkpoint, but since we're not
	// verifying it we can use a dummy one.
	// This seems to be the minimal way to create a verifier that
	// the parser will accept.
	dummyVerifier, err := tdbnote.NewVerifier("dummy\n\n")
	if err != nil {
		klog.Exitf("Failed to create dummy verifier: %v", err)
	}
	cp, _, _, err := log.ParseCheckpoint(cpRaw, "dummy", note.VerifierList(dummyVerifier))
	if err != nil {
		klog.Exitf("Failed to parse checkpoint: %v", err)
	}

	if li >= cp.Size {
		klog.Exitf("Leaf index %d is out of range for log size %d", li, cp.Size)
	}

	bundleIndex := li / uint64(layout.EntryBundleWidth)
	indexInBundle := li % uint64(layout.EntryBundleWidth)

	bundle, err := client.GetEntryBundle(ctx, fetcher.EntryBundle, bundleIndex, cp.Size)
	if err != nil {
		klog.Exitf("Failed to get entry bundle: %v", err)
	}

	if int(indexInBundle) >= len(bundle.Entries) {
		klog.Exitf("Index %d is out of range for bundle of size %d", indexInBundle, len(bundle.Entries))
	}
	entryData := bundle.Entries[indexInBundle]

	var entry staticct.Entry
	if err := entry.UnmarshalText(entryData); err != nil {
		klog.Exitf("Failed to unmarshal entry: %v", err)
	}

	certBytes := entry.Certificate
	if entry.IsPrecert {
		// For precertificates, the `Certificate` field holds the TBSCertificate.
		// We need to wrap this in a `Certificate` structure to be able to parse it.
		// This is a bit of a hack, but it's what the `x509` package expects.
		cert, err := x509.ParseCertificate(entry.Precertificate)
		if err != nil {
			klog.Exitf("Failed to parse precertificate: %v", err)
		}
		certBytes = cert.Raw
	}

	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}

	if err := pem.Encode(os.Stdout, pemBlock); err != nil {
		klog.Exitf("Failed to encode PEM: %v", err)
	}
}
