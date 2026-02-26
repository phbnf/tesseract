package ct

import (
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create opts with small max body size
	opts := hOpts()
	opts.MaxBodySize = 10 // 10 bytes

	// Use NewPathHandlers to setup the handlers
	handlers := NewPathHandlers(t.Context(), opts, log)

	// Get the handler for add-chain
	addChainPath := path.Join(prefix, rfc6962.AddChainPath)
	handler, ok := handlers[addChainPath]
	if !ok {
		t.Fatalf("Handler not found: %s", addChainPath)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	// Send a request larger than 10 bytes
	body := strings.Repeat("a", 11)
	resp, err := http.Post(server.URL+addChainPath, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status 413, got %d", resp.StatusCode)
	}
}

func TestMaxBodySizePreChain(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create opts with small max body size
	opts := hOpts()
	opts.MaxBodySize = 10 // 10 bytes

	// Use NewPathHandlers to setup the handlers
	handlers := NewPathHandlers(t.Context(), opts, log)

	// Get the handler for add-pre-chain
	addPreChainPath := path.Join(prefix, rfc6962.AddPreChainPath)
	handler, ok := handlers[addPreChainPath]
	if !ok {
		t.Fatalf("Handler not found: %s", addPreChainPath)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	// Send a request larger than 10 bytes
	body := strings.Repeat("a", 11)
	resp, err := http.Post(server.URL+addPreChainPath, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status 413, got %d", resp.StatusCode)
	}
}
