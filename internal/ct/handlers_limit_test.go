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
	"bytes"
	"net/http"
	"path"
	"testing"

	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestAddChainBodyLimit(t *testing.T) {
	log, _ := setupTestLog(t)
	// We need a larger deadline for large body if strict checking is on, but it shouldn't matter for the error.
	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	// 5MB body, exceeding the 4MB limit
	size := 5 * 1024 * 1024
	body := make([]byte, size)
	// Fill with spaces to mimic potential whitespace padding
	for i := range body {
		body[i] = ' '
	}

	resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to post: %v", err)
	}
	defer resp.Body.Close()

	if got, want := resp.StatusCode, http.StatusRequestEntityTooLarge; got != want {
		t.Errorf("StatusCode=%d, want %d", got, want)
	}
}
