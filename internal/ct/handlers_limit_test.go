package ct

import (
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)
	hhOpts := hOpts()
	hhOpts.MaxBodySize = 10 // Very small limit

	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hhOpts)
	defer server.Close()

	body := strings.Repeat("a", 20) // Larger than 10
	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", rfc6962.AddChainPath, err)
	}
	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", rfc6962.AddChainPath, got, want)
	}

	// Also test a small request passes (well, it will fail parsing, but not with 413)
	hhOpts.MaxBodySize = 1024
	server2 := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hhOpts)
	defer server2.Close()

	resp2, err := http.Post(server2.URL+rfc6962.AddChainPath, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", rfc6962.AddChainPath, err)
	}
	// {} is invalid chain, so it should return 400, NOT 413.
	if got, want := resp2.StatusCode, http.StatusBadRequest; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", rfc6962.AddChainPath, got, want)
	}
}
