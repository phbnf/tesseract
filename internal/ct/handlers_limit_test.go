package ct

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainBodyLimit(t *testing.T) {
	log, _ := setupTestLog(t)
	handlers := NewPathHandlers(t.Context(), hOpts(), log)
	handler := handlers[path.Join(prefix, rfc6962.AddChainPath)]
	s := httptest.NewServer(handler)
	defer s.Close()

	// 5MB body of ' ' (spaces) which is valid JSON whitespace until the end?
	// If we just use zeros, it might fail JSON parsing earlier if it peeks.
	// But io.ReadAll reads everything.
	// Let's use spaces.
	largeBody := make([]byte, 5*1024*1024)
	for i := range largeBody {
		largeBody[i] = ' '
	}

	resp, err := http.Post(s.URL+rfc6962.AddChainPath, "application/json", bytes.NewReader(largeBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected 413 Request Entity Too Large, got %d", resp.StatusCode)
	}
}
