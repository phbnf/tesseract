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
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandlersRejectLargeBody(t *testing.T) {
	log, _ := setupTestLog(t)
	handlers := NewPathHandlers(context.Background(), hOpts(), log)

	// Create a body slightly larger than the expected limit (e.g. 4MB + 1 byte)
	// We'll use a limit of 4MB.
	const maxBodySize = 4 * 1024 * 1024
	largeBody := strings.Repeat("a", maxBodySize+1)

	for path, handler := range postHandlers(t, handlers) {
		t.Run(path, func(t *testing.T) {
			// Wrap the handler to mimic what ctlog.go does.
			wrappedHandler := http.MaxBytesHandler(handler, int64(MaxBodySize))
			s := httptest.NewServer(wrappedHandler)
			defer s.Close()

			resp, err := http.Post(s.URL+path, "application/json", strings.NewReader(largeBody))
			if err != nil {
				t.Fatalf("http.Post(%s)=(_,%q); want (_,nil)", path, err)
			}

			if resp.StatusCode != http.StatusRequestEntityTooLarge {
				t.Errorf("http.Post(%s) status code = %d; want %d (Request Entity Too Large)", path, resp.StatusCode, http.StatusRequestEntityTooLarge)
			}
		})
	}
}
