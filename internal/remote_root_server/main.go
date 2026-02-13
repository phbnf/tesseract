package main

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"github.com/transparency-dev/tesseract/internal/testdata"
)

var (
	httpEndpoint   = flag.String("http_endpoint", ":8080", "The endpoint to run the server on")
	tesseractURL   = flag.String("tesseract_url", "http://localhost:6962", "Base URL of the Tesseract server to verify")
	rootsBackupDir = flag.String("roots_backup_dir", "", "Directory where Tesseract backs up roots (optional)")
	verifyInterval = flag.Duration("verify_interval", 5*time.Second, "Interval between verification attempts")
	rootsReject    = flag.String("roots_reject_fingerprints", "", "Comma-separated list of SHA256 fingerprints to reject")
	exitOnSuccess  = flag.Bool("exit_on_success", false, "Exit with code 0 after successful verification")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// Generate CSV content
	csvContent, fingerprints, err := generateRootsCSV()
	if err != nil {
		klog.Errorf("Failed to generate roots CSV: %v", err)
		os.Exit(1)
	}

	http.HandleFunc("/roots.csv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		if _, err := w.Write([]byte(csvContent)); err != nil {
			klog.Warningf("Failed to serve roots.csv: %v", err)
		}
		klog.Infof("Served roots.csv: remote_addr=%v", r.RemoteAddr)
	})

	rejected := make(map[string]struct{})
	if *rootsReject != "" {
		for _, fps := range strings.Split(*rootsReject, ",") {
			fp := strings.TrimSpace(fps)
			if fp != "" {
				rejected[strings.ToLower(fp)] = struct{}{}
			}
		}
	}

	// Start verification in background
	go verifyLoop(fingerprints, rejected)

	klog.Infof("Starting remote_root_server: endpoint=%v", *httpEndpoint)
	if err := http.ListenAndServe(*httpEndpoint, nil); err != nil {
		klog.Errorf("Server failed: %v", err)
		os.Exit(1)
	}
}

func generateRootsCSV() (string, []string, error) {
	certs := []string{testdata.CACertPEM, testdata.FakeCACertPEM}
	var fingerprints []string
	var b strings.Builder
	w := csv.NewWriter(&b)

	// Write header
	if err := w.Write([]string{"Subject", "CA Owner", "X.509 Certificate (PEM)", "SHA-256 Fingerprint", "Intended Use Case(s) Served"}); err != nil {
		return "", nil, err
	}

	for _, pemData := range certs {
		block, _ := pem.Decode([]byte(pemData))
		if block == nil {
			return "", nil, fmt.Errorf("failed to decode PEM")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return "", nil, fmt.Errorf("failed to parse certificate: %v", err)
		}

		fingerprint := sha256.Sum256(cert.Raw)
		fingerprintHex := strings.ToLower(hex.EncodeToString(fingerprint[:]))
		fingerprints = append(fingerprints, fingerprintHex)

		subject := cert.Subject.String()
		caOwner := subject

		record := []string{
			subject,
			caOwner,
			// Make sure we write the exact PEM content without extra modifications if possible,
			// but csv writer will handle quoting.
			pemData,
			fingerprintHex,
			"Server Authentication (TLS) 1.3.6.1.5.5.7.3.1",
		}
		if err := w.Write(record); err != nil {
			return "", nil, err
		}
	}
	w.Flush()
	return b.String(), fingerprints, nil
}

func verifyLoop(fingerprints []string, rejected map[string]struct{}) {
	if *exitOnSuccess {
		// Run immediately once
		if err := verifyRoots(fingerprints, rejected); err != nil {
			klog.Errorf("Verification FAILED: %v", err)
		} else {
			klog.Infof("Verification PASSED. Exiting.")
			os.Exit(0)
		}
	}

	ticker := time.NewTicker(*verifyInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := verifyRoots(fingerprints, rejected); err != nil {
			klog.Errorf("Verification FAILED: %v", err)
		} else {
			klog.V(1).Infof("Verification PASSED")
			if *exitOnSuccess {
				klog.Infof("Exiting after successful verification.")
				os.Exit(0)
			}
		}
	}
}

func verifyRoots(fingerprints []string, rejected map[string]struct{}) error {
	// 1. Check get-roots
	url, err := url.JoinPath(*tesseractURL, "/ct/v1/get-roots")
	if err != nil {
		return fmt.Errorf("can't build get-root url: %v", err)
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to call get-roots: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			klog.Warningf("resp.Body.Close(): %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get-roots returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	var parsed struct {
		Certificates [][]byte `json:"certificates"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	foundFingerprints := make(map[string]struct{})
	for _, certDER := range parsed.Certificates {
		fingerprint := sha256.Sum256(certDER)
		fp := strings.ToLower(hex.EncodeToString(fingerprint[:]))
		foundFingerprints[fp] = struct{}{}

		if _, ok := rejected[fp]; ok {
			return fmt.Errorf("found rejected root: %s", fp)
		}
	}

	// Verify all expected fingerprints are present (unless rejected)
	for _, fp := range fingerprints {
		_, isRejected := rejected[fp]
		_, isFound := foundFingerprints[fp]

		if isRejected {
			if isFound {
				return fmt.Errorf("expected rejected root %s to be absent, but it was present", fp)
			}
			continue
		}
		if !isFound {
			return fmt.Errorf("missing expected root: %s", fp)
		}
		klog.V(1).Infof("Found expected root in get-roots's response: fingerprint=%v", fp)
	}

	// 2. Check backup dir (if configured)
	if *rootsBackupDir != "" {
		files, err := os.ReadDir(*rootsBackupDir)
		if err != nil {
			return fmt.Errorf("failed to read roots backup dir: %v", err)
		}

		if len(files) == 0 {
			return fmt.Errorf("no roots found in backup dir %s", *rootsBackupDir)
		}

		// Prepare map of expected roots that should be in backup
		// Even if a cert is rejected, it should be in the backup
		notInBackupFingerprints := make(map[string]struct{})
		for _, fp := range fingerprints {
			notInBackupFingerprints[fp] = struct{}{}
		}

		for _, file := range files {
			if file.IsDir() {
				// We expect a flat directory of roots
				continue
			}

			// User requested to trust the filename as the fingerprint hash
			fp := strings.ToLower(file.Name())

			if _, ok := notInBackupFingerprints[fp]; ok {
				delete(notInBackupFingerprints, fp)
				klog.V(1).Infof("Found valid root in backup: fingerprint=%v", fp)
			} else {
				// If it's not in notInBackupFingerprints, it's either an unexpected root
				// or a duplicate (not possible in a single directory).
				return fmt.Errorf("found unexpected root in backup: %s", fp)
			}
		}

		if len(notInBackupFingerprints) > 0 {
			var missing []string
			for fp := range notInBackupFingerprints {
				missing = append(missing, fp)
			}
			return fmt.Errorf("missing expected roots in backup: %v", missing)
		}
	}

	return nil
}
