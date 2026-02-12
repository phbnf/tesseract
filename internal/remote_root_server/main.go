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
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/transparency-dev/tesseract/internal/testdata"
)

var (
	httpEndpoint    = flag.String("http_endpoint", ":8080", "The endpoint to run the server on")
	tesseractURL    = flag.String("tesseract_url", "http://localhost:6962", "Base URL of the Tesseract server to verify")
	rootsBackupDir  = flag.String("roots_backup_dir", "", "Directory where Tesseract backs up roots (optional)")
	verifyInterval  = flag.Duration("verify_interval", 5*time.Second, "Interval between verification attempts")
	rootsReject     = flag.String("roots_reject_fingerprints", "", "Comma-separated list of SHA256 fingerprints to reject")
	exitOnSuccess   = flag.Bool("exit_on_success", false, "Exit with code 0 after successful verification")
)

func main() {
	flag.Parse()

	// Generate CSV content
	csvContent, expectedFingerprints, err := generateRootsCSV()
	if err != nil {
		log.Fatalf("Failed to generate roots CSV: %v", err)
	}

	http.HandleFunc("/roots.csv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		if _, err := w.Write([]byte(csvContent)); err != nil {
			log.Printf("Failed to serve roots.csv: %v", err)
		}
		log.Printf("Served roots.csv to %s", r.RemoteAddr)
	})

	rejected := make(map[string]bool)
	if *rootsReject != "" {
		for _, fps := range strings.Split(*rootsReject, ",") {
			fp := strings.TrimSpace(fps)
			if fp != "" {
				rejected[strings.ToUpper(fp)] = true
			}
		}
	}

	// Start verification in background
	go verifyLoop(expectedFingerprints, rejected)

	log.Printf("Starting remote_root_server on %s", *httpEndpoint)
	if err := http.ListenAndServe(*httpEndpoint, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
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
		fingerprintHex := strings.ToUpper(hex.EncodeToString(fingerprint[:]))
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

func verifyLoop(expectedFingerprints []string, rejected map[string]bool) {
	if *exitOnSuccess {
		// Run immediately once
		if err := verifyRoots(expectedFingerprints, rejected); err != nil {
			log.Printf("Verification FAILED: %v", err)
		} else {
			log.Printf("Verification PASSED. Exiting.")
			os.Exit(0)
		}
	}

	ticker := time.NewTicker(*verifyInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := verifyRoots(expectedFingerprints, rejected); err != nil {
			log.Printf("Verification FAILED: %v", err)
		} else {
			log.Printf("Verification PASSED")
			if *exitOnSuccess {
				log.Printf("Exiting after successful verification.")
				os.Exit(0)
			}
		}
	}
}

func verifyRoots(expectedFingerprints []string, rejected map[string]bool) error {
	// 1. Check get-roots
	resp, err := http.Get(*tesseractURL + "/ct/v1/get-roots")
	if err != nil {
		return fmt.Errorf("failed to call get-roots: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("resp.Body.Close(): %v", err)
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

	foundFingerprints := make(map[string]bool)
	for _, certDER := range parsed.Certificates {
		fingerprint := sha256.Sum256(certDER)
		fp := strings.ToUpper(hex.EncodeToString(fingerprint[:]))
		foundFingerprints[fp] = true
		
		if rejected[fp] {
			return fmt.Errorf("found rejected root: %s", fp)
		}
	}

	// Verify all expected fingerprints are present (unless rejected)
	for _, fp := range expectedFingerprints {
		if rejected[fp] {
			if foundFingerprints[fp] {
				return fmt.Errorf("expected rejected root %s to be absent, but it was present", fp)
			}
			continue
		}
		if !foundFingerprints[fp] {
			return fmt.Errorf("missing expected root: %s", fp)
		}
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
	}

	return nil
}
