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
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// GetFetcher returns an S3 read function for objects in a given bucket.
func GetFetcher(ctx context.Context, projectID string, bucket string) (func(ctx context.Context, path string) ([]byte, error), error) {
	// TODO(phboneff): this should probably move somewhere else
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default AWS configuration: %v", err)
	}
	c := s3.NewFromConfig(sdkConfig)

	return func(ctx context.Context, path string) ([]byte, error) {
		r, err := c.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path),
		})

		if err != nil {
			return nil, fmt.Errorf("getObject: failed to create reader for object %q in bucket %q: %w", path, bucket, err)
		}

		d, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read %q: %v", path, err)
		}
		return d, r.Body.Close()
	}, nil
}
