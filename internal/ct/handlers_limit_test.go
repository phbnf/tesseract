// Copyright 2025 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	// Create options with a small body size limit.
	opts := hOpts()
	opts.MaxBodySize = 10 // Very small limit

	handlers := NewPathHandlers(t.Context(), opts, log)
	addChainPath := path.Join(prefix, rfc6962.AddChainPath)
	handler, ok := handlers[addChainPath]
	if !ok {
		t.Fatalf("Handler not found: %s", addChainPath)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	// Create a body larger than the limit (11 bytes)
	body := strings.Repeat("a", 11)
	resp, err := http.Post(server.URL, "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", addChainPath, err)
	}

	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", addChainPath, got, want)
	}
}
