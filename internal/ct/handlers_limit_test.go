// Copyright 2024 Google LLC. All Rights Reserved.
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
	"bytes"
	"net/http"
	"path"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainTooLarge(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create a large body > 4MB (MaxBodySize)
	// We use 5MB to be safe.
	largeBody := make([]byte, 5*1024*1024)
	// Fill with something that looks like start of JSON to avoid immediate failure if possible,
	// though MaxBytesReader should fail during Read regardless of content.
	copy(largeBody, []byte(`{"chain": [`))

	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", bytes.NewReader(largeBody))
	if err != nil {
		t.Fatalf("http.Post: %v", err)
	}
	defer resp.Body.Close()

	// We expect 413 Payload Too Large
	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("http.Post(%s)=(%d,nil); want (%d,nil)", rfc6962.AddChainPath, got, want)
	}
}
