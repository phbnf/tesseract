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
	"net/http/httptest"
	"path"
	"testing"
	"github.com/transparency-dev/tesseract/internal/types/rfc6962"
)

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create handlers with explicit small limit for testing
	opts := hOpts()
	opts.MaxBodySize = 1024 // 1KB
	handlers := NewPathHandlers(t.Context(), opts, log)
	server := httptest.NewServer(handlers[path.Join(prefix, rfc6962.AddChainPath)])
	defer server.Close()

	tests := []struct {
		name       string
		bodySize   int
		expectCode int
	}{
		{
			name:       "small-body",
			bodySize:   100,
			expectCode: http.StatusBadRequest, // Invalid JSON, but size is OK
		},
		{
			name:       "exact-limit",
			bodySize:   1024,
			expectCode: http.StatusBadRequest, // Invalid JSON, but size is OK
		},
		{
			name:       "too-large-body",
			bodySize:   1025,
			expectCode: http.StatusRequestEntityTooLarge,
		},
		{
			name:       "way-too-large-body",
			bodySize:   10 * 1024 * 1024,
			expectCode: http.StatusRequestEntityTooLarge,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body := make([]byte, test.bodySize)
			// Fill with 'a' to be valid utf-8 but invalid JSON
			for i := range body {
				body[i] = 'a'
			}
			resp, err := http.Post(server.URL+rfc6962.AddChainPath, "application/json", bytes.NewReader(body))
			if err != nil {
				t.Fatalf("Failed to post: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != test.expectCode {
				t.Errorf("Got status %d, expected %d", resp.StatusCode, test.expectCode)
			}
		})
	}
}
