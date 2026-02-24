package ct

import (
	"net/http"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)

	// Create options with a small limit for testing
	opts := hOpts()
	opts.MaxBodySize = 10 // Small limit (10 bytes)

	// Path construction must match how setupTestServer works.
	// In handlers_test.go, setupTestServer uses handlers[path].
	// NewPathHandlers constructs paths as prefix + rfc6962 path.
	path := prefix + rfc6962.AddChainPath
	server := setupTestServer(t, log, path, opts)
	defer server.Close()

	// Test Case 1: Body larger than limit
	t.Run("BodyTooLarge", func(t *testing.T) {
		body := strings.Repeat("a", 20)
		resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if resp.StatusCode != http.StatusRequestEntityTooLarge {
			t.Errorf("Expected status 413, got %d", resp.StatusCode)
		}
	})

	// Test Case 2: Body within limit
	// Note: It will likely fail JSON parsing or chain validation, but it should NOT be 413.
	t.Run("BodyWithinLimit", func(t *testing.T) {
		body := strings.Repeat("a", 5)
		resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if resp.StatusCode == http.StatusRequestEntityTooLarge {
			t.Errorf("Did not expect status 413, got %d", resp.StatusCode)
		}
		// Expect 400 Bad Request because "aaaaa" is not valid JSON
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", resp.StatusCode)
		}
	})

	// Test Case 3: Default Limit (implicitly 0 -> 4MB)
	t.Run("DefaultLimit", func(t *testing.T) {
		optsDefault := hOpts() // MaxBodySize is 0
		// We need a separate server because opts are baked into the handler closure
		// Note: setupTestServer creates a new handler instance using the passed opts.
		serverDefault := setupTestServer(t, log, path, optsDefault)
		defer serverDefault.Close()

		// Should accept small body (well within 4MB)
		body := strings.Repeat("a", 100)
		resp, err := http.Post(serverDefault.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if resp.StatusCode == http.StatusRequestEntityTooLarge {
			t.Errorf("Did not expect status 413 for default limit with small body, got %d", resp.StatusCode)
		}
	})

	// Test Case 4: PreChain Limit
	t.Run("PreChainLimit", func(t *testing.T) {
		opts := hOpts()
		opts.MaxBodySize = 10
		path := prefix + rfc6962.AddPreChainPath
		server := setupTestServer(t, log, path, opts)
		defer server.Close()

		body := strings.Repeat("a", 20)
		resp, err := http.Post(server.URL+rfc6962.AddPreChainPath, "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if resp.StatusCode != http.StatusRequestEntityTooLarge {
			t.Errorf("Expected status 413, got %d", resp.StatusCode)
		}
	})
}
