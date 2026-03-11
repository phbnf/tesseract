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
	"path"
	"strings"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChain_RequestTooLarge(t *testing.T) {
	log, _ := setupTestLog(t)
	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	// Create a large body slightly larger than MaxBodySize
	// MaxBodySize is 4MB. We construct a JSON body that exceeds it.
	// {"chain": ["..."]}
	// We need enough padding.
	padding := strings.Repeat("a", MaxBodySize)
	jsonBody := `{"chain": ["` + padding + `"]}`

	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", strings.NewReader(jsonBody))
	if err != nil {
		t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", rfc6962.AddChainPath, err)
	}
	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", rfc6962.AddChainPath, got, want)
	}
}
