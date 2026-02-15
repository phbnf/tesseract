package ct

import (
	"net/http"
	"path"
	"strings"
	"testing"
)

func TestAddChainLargeBody(t *testing.T) {
	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, path.Join(prefix, "ct/v1/add-chain"), hOpts())
	defer server.Close()

	// Create a body larger than the proposed limit (e.g. > 4MB)
	largeBody := strings.Repeat("a", 5*1024*1024)
	reqBody := `{"chain": ["` + largeBody + `"]}`

	resp, err := http.Post(server.URL+"/ct/v1/add-chain", "application/json", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status 413 (Request Entity Too Large), got %d", resp.StatusCode)
	}
}
