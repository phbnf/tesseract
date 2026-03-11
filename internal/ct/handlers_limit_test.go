package ct

import (
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainLargeBody(t *testing.T) {
	log, _ := setupTestLog(t)
	handlers := NewPathHandlers(t.Context(), hOpts(), log)

	// Create a body larger than what we intend to limit (e.g. > 4MB)
	// 5MB string
	largeBody := strings.Repeat("a", 5*1024*1024)

	handler := handlers[path.Join(prefix, rfc6962.AddChainPath)]
	s := httptest.NewServer(handler)
	defer s.Close()

	resp, err := http.Post(s.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(largeBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}

	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
	}
}
