// Copyright 2025 The Tessera authors. All Rights Reserved.
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

package inmemory

import (
	"context"

	"github.com/transparency-dev/static-ct/storage"
	"k8s.io/klog/v2"
)

// IssuersStorage is a map to store of Issuers bytes keys on a string, like its hash.
type IssuersStorage map[string][]byte

// NewIssuerStorage creates a new IssuerStorage.
func NewIssuerStorage() (*IssuersStorage, error) {
	m := IssuersStorage(make(map[string][]byte))
	return &m, nil
}

// AddIssuers stores Issuers values under their Key if there isn't an object under Key already.
func (s IssuersStorage) AddIssuersIfNotExist(_ context.Context, kv []storage.KV) error {
	for _, kv := range kv {
		objName := string(kv.K)
		if _, ok := s[objName]; ok {
			klog.V(2).Infof("AddIssuersIfNotExist: object %q already exists, continuing", objName)
			continue
		}
		s[objName] = kv.V
	}
	return nil
}
