// Copyright 2024 The Tessera authors. All Rights Reserved.
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

package aws

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/smithy-go"
	sctfe "github.com/transparency-dev/static-ct"
	"k8s.io/klog/v2"
)

// IssuersStorage is a key value store backed by S3 on AWS to store issuer chains.
type IssuersStorage struct {
	s3Client    *s3.Client
	bucket      string
	prefix      string
	contentType string
}

// NewIssuerStorage creates a new IssuerStorage.
//
// The specified bucket must exist or an error will be returned.
func NewIssuerStorage(ctx context.Context, projectID string, bucket string, prefix string, contentType string) (*IssuersStorage, error) {
	// TODO(phboneff): this should probably move somewhere else
	// TODO(phboneff): withjsonreads?
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default AWS configuration: %v", err)
	}
	c := s3.NewFromConfig(sdkConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	r := &IssuersStorage{
		s3Client:    c,
		bucket:      bucket,
		prefix:      prefix,
		contentType: contentType,
	}

	return r, nil
}

// keyToObjName converts bytes to a GCS object name.
func (s *IssuersStorage) keyToObjName(key []byte) string {
	return path.Join(s.prefix, string(key))
}

// AddIssuers stores Issuers values under their Key if there isn't an object under Key already.
func (s *IssuersStorage) AddIssuersIfNotExist(ctx context.Context, kv []sctfe.KV) error {
	// We first try and see if this issuer cert has already been stored since reads
	// are cheaper than writes.
	// TODO(phboneff): add parallel operations
	for _, kv := range kv {
		objName := s.keyToObjName(kv.K)
		put := &s3.PutObjectInput{
			Bucket:      aws.String(s.bucket),
			Key:         aws.String(objName),
			Body:        bytes.NewReader(kv.V),
			ContentType: aws.String(s.contentType),
			IfNoneMatch: aws.String("*"),
		}

		if _, err := s.s3Client.PutObject(ctx, put); err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr); apiErr.ErrorCode() == "PreconditionFailed" {
				klog.V(2).Infof("AddIssuersIfNotExist: object %q already exists in bucket %q, continuing", objName, s.bucket)
				return nil
			}
			return fmt.Errorf("failed to write object %q to bucket %q: %w", objName, s.bucket, err)
		}
		klog.V(2).Infof("AddIssuersIfNotExist: added %q in bucket %q", objName, s.bucket)
	}
	return nil
}
