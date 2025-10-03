// Copyright 2016 Google LLC. All Rights Reserved.
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

package tesseract

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewCertValidationOpts(t *testing.T) {
	t100 := time.Unix(100, 0)
	t200 := time.Unix(200, 0)

	for _, tc := range []struct {
		desc    string
		wantErr string
		cvCfg   ChainValidationConfig
	}{
		{
			desc:    "empty-rootsPemFile",
			wantErr: "empty rootsPemFile",
		},
		{
			desc:    "missing-root-cert",
			wantErr: "failed to read trusted roots",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/bogus.cert",
			},
		},
		{
			desc:    "rejecting-all",
			wantErr: "configuration would reject all certificates",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:    "./internal/testdata/fake-ca.cert",
				RejectExpired:   true,
				RejectUnexpired: true},
		},
		{
			desc:    "unknown-ext-key-usage-1",
			wantErr: "unknown extended key usage",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/fake-ca.cert",
				ExtKeyUsages: "wrong_usage"},
		},
		{
			desc:    "unknown-ext-key-usage-2",
			wantErr: "unknown extended key usage",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/fake-ca.cert",
				ExtKeyUsages: "ClientAuth,ServerAuth,TimeStomping",
			},
		},
		{
			desc:    "unknown-ext-key-usage-3",
			wantErr: "unknown extended key usage",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/fake-ca.cert",
				ExtKeyUsages: "Any ",
			},
		},
		{
			desc:    "unknown-reject-ext",
			wantErr: "failed to parse RejectExtensions",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:     "./internal/testdata/fake-ca.cert",
				RejectExtensions: "1.2.3.4,one.banana.two.bananas",
			},
		},
		{
			desc:    "limit-before-start",
			wantErr: "before start",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:  "./internal/testdata/fake-ca.cert",
				NotAfterStart: &t200,
				NotAfterLimit: &t100,
			},
		},
		{
			desc: "ok",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/fake-ca.cert",
			},
		},
		{
			desc: "ok-ext-key-usages",
			cvCfg: ChainValidationConfig{
				RootsPEMFile: "./internal/testdata/fake-ca.cert",
				ExtKeyUsages: "ServerAuth,ClientAuth,OCSPSigning",
			},
		},
		{
			desc: "ok-reject-ext",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:     "./internal/testdata/fake-ca.cert",
				RejectExtensions: "1.2.3.4,5.6.7.8",
			},
		},
		{
			desc: "ok-start-timestamp",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:  "./internal/testdata/fake-ca.cert",
				NotAfterStart: &t100,
			},
		},
		{
			desc: "ok-limit-timestamp",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:  "./internal/testdata/fake-ca.cert",
				NotAfterStart: &t200,
			},
		},
		{
			desc: "ok-range-timestamp",
			cvCfg: ChainValidationConfig{
				RootsPEMFile:  "./internal/testdata/fake-ca.cert",
				NotAfterStart: &t100,
				NotAfterLimit: &t200,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			vc, err := newChainValidator(tc.cvCfg)
			if len(tc.wantErr) == 0 && err != nil {
				t.Errorf("ValidateLogConfig()=%v, want nil", err)
			}
			if len(tc.wantErr) > 0 && (err == nil || !strings.Contains(err.Error(), tc.wantErr)) {
				t.Errorf("ValidateLogConfig()=%v, want err containing %q", err, tc.wantErr)
			}
			if err == nil && vc == nil {
				t.Error("err and ValidatedLogConfig are both nil")
			}
		})
	}
}

func TestNotBeforeRLFromFlagValue(t *testing.T) {
	tests := []struct {
		name      string
		flagValue string
		want      *NotBeforeRL
		wantErr   bool
	}{
		{
			name:      "ok",
			flagValue: "10m:10",
			want: &NotBeforeRL{
				AgeThreshold: 10 * time.Minute,
				RateLimit:    10,
			},
			wantErr: false,
		},
		{
			name:      "wrong-format",
			flagValue: "10m10",
			wantErr:   true,
		},
		{
			name:      "not-a-duration",
			flagValue: "10:10",
			wantErr:   true,
		},
		{
			name:      "not-a-rate",
			flagValue: "10:a",
			wantErr:   true,
		},
		{
			name:      "null-duration",
			flagValue: "0m:10",
			wantErr:   true,
		},
		{
			name:      "negative-duration",
			flagValue: "-10m:10",
			wantErr:   true,
		},
		{
			name:      "negative-rate",
			flagValue: "10m:-10",
			wantErr:   true,
		},
		{
			name:      "null-rate",
			flagValue: "10m:0",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NotBeforeRLFromFlagValue(tt.flagValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotBeforeRLFromFlagValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotBeforeRLFromFlagValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
