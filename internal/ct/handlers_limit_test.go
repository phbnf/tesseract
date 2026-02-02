package ct

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainLargeBody(t *testing.T) {
	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	// Create a body larger than the proposed 4MB limit
	largeBody := bytes.Repeat([]byte("a"), 5*1024*1024) // 5MB

    // We construct a JSON string manually to avoid loading everything into memory
    // before the test (though bytes.Repeat does that anyway).
    // The point is to send a large payload.
    // Use a reader to stream it.

    // Actually, simple way:
    // Create a reader from the large bytes.
    hugeJSON := `{"chain": ["` + string(largeBody) + `"]}`

    resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(hugeJSON))
    if err != nil {
        t.Fatalf("http.Post failed: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusRequestEntityTooLarge {
        t.Errorf("Expected 413 RequestEntityTooLarge, got %v", resp.StatusCode)
    }
}
