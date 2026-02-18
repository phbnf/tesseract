// Copyright 2025 The Tessera Authors. All Rights Reserved.
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

func TestAddChainLargeBody(t *testing.T) {
	log, _ := setupTestLog(t)
	// We want to test that a large body is rejected if we set a limit.
	// Since we haven't implemented the limit yet, this test serves as a placeholder
	// to verify I can run tests and later to verify the fix.

	// Create options with MaxBodySize set to a small limit.
	opts := hOpts()
	opts.MaxBodySize = 10 // 10 bytes limit

	handlers := NewPathHandlers(t.Context(), opts, log)
	pathStr := path.Join(prefix, rfc6962.AddChainPath)
	handler, ok := handlers[pathStr]
	if !ok {
		t.Fatalf("Handler not found: %s", pathStr)
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	// 1KB body, which is much larger than 10 bytes.
	largeBody := strings.Repeat("a", 1024)

	resp, err := http.Post(server.URL+pathStr, "application/json", strings.NewReader(largeBody))
	if err != nil {
		t.Fatalf("http.Post failed: %v", err)
	}
	defer resp.Body.Close()

	// It should return 413 Payload Too Large.
	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected 413 Payload Too Large, got %d", resp.StatusCode)
	}
}
