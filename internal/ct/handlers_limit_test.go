package ct

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAddChain_BodySizeLimit(t *testing.T) {
	// Create a body larger than the intended limit (e.g., > 4MB)
	largeBody := strings.Repeat("a", 5*1024*1024)
    reqBody := `{"chain": ["` + largeBody + `"]}`

	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, prefix + "/ct/v1/add-chain", hOpts())
	defer server.Close()

	resp, err := http.Post(server.URL+prefix+"/ct/v1/add-chain", "application/json", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}

    // It should fail with 413 Payload Too Large (or Request Entity Too Large)
	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected 413 Request Entity Too Large, got %d", resp.StatusCode)
	}

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }
    bodyStr := string(bodyBytes)

    // We expect the error to mention the body size limit being exceeded
    if !strings.Contains(bodyStr, "request body too large") {
        t.Errorf("Expected error 'request body too large', got: %s", bodyStr)
    }
}
