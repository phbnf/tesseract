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

func TestMaxBodySize(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create a large body (e.g., 5MB)
	largeBody := make([]byte, 5*1024*1024)

	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddChainPath), hOpts())
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+rfc6962.AddChainPath, bytes.NewReader(largeBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status code 413, got %v", resp.StatusCode)
	}
}

func TestMaxBodySizePreChain(t *testing.T) {
	log, _ := setupTestLog(t)
	// Create a large body (e.g., 5MB)
	largeBody := make([]byte, 5*1024*1024)

	server := setupTestServer(t, log, path.Join(prefix, rfc6962.AddPreChainPath), hOpts())
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL+rfc6962.AddPreChainPath, bytes.NewReader(largeBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("Expected status code 413, got %v", resp.StatusCode)
	}
}
