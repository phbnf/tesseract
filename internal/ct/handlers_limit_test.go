package ct

import (
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainRequestLimit(t *testing.T) {
	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	// Create a body larger than the expected limit (e.g., 5MB)
	// We use a valid JSON structure but with a large ignored field
	largeBody := `{ "chain": [], "garbage": "` + strings.Repeat("a", 5*1024*1024) + `" }`

	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(largeBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}
	defer resp.Body.Close()

	// Expect 413 Payload Too Large
	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status 413 Payload Too Large, got %d", resp.StatusCode)
	}
}
