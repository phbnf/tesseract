package ct

import (
	"bytes"
	"net/http"
	"path"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainBodyLimit(t *testing.T) {
	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

    // Case 1: Large body (> 4MB) should fail with 413
	largeBody := make([]byte, 5*1024*1024)
    copy(largeBody, []byte(`{"chain": ["`))
	req, err := http.NewRequest("POST", server.URL+rfc6962.AddChainPath, bytes.NewReader(largeBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
    req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

    if resp.StatusCode != http.StatusRequestEntityTooLarge {
        t.Errorf("Large body: expected status 413, got %d", resp.StatusCode)
    }

    // Case 2: Small body (< 4MB) should NOT fail with 413 (it might fail with 400 because of invalid JSON or empty chain, which is fine)
    smallBody := []byte(`{"chain": []}`)
    reqSmall, err := http.NewRequest("POST", server.URL+rfc6962.AddChainPath, bytes.NewReader(smallBody))
    if err != nil {
        t.Fatalf("Failed to create small request: %v", err)
    }
    reqSmall.Header.Set("Content-Type", "application/json")

    respSmall, err := client.Do(reqSmall)
    if err != nil {
        t.Fatalf("Failed to perform small request: %v", err)
    }
    defer respSmall.Body.Close()

    if respSmall.StatusCode == http.StatusRequestEntityTooLarge {
        t.Errorf("Small body: unexpected status 413")
    }
}
