package ct

import (
	"bytes"
	"net/http"
	"path"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create handler options with a small MaxBodySize
	opts := hOpts()
	opts.MaxBodySize = 10 // 10 bytes limit

	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), opts)
	defer server.Close()

	// Create a body larger than 10 bytes
	body := bytes.NewBuffer(make([]byte, 20))
	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", body)
	if err != nil {
		t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", rfc6962.AddChainPath, err)
	}

	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", rfc6962.AddChainPath, got, want)
	}
}
