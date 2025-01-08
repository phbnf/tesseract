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

package sctfe

import (
	"crypto"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/certificate-transparency-go/asn1"
	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/certificate-transparency-go/x509util"
	"k8s.io/klog/v2"
)

type CertValidationConfig struct {
	RootsPemFile     string
	RejectExpired    bool
	RejectUnexpired  bool
	ExtKeyUsages     string
	RejectExtensions string
	NotAfterStart    *time.Time
	NotAfterLimit    *time.Time
}

// ValidatedLogConfig represents the LogConfig with the information that has
// been successfully parsed as a result of validating it.
type ValidatedLogConfig struct {
	// Origin identifies the log. It will be used in its checkpoint, and
	// is also its submission prefix, as per https://c2sp.org/static-ct-api.
	Origin string
	// Used to sign the checkpoint and SCTs.
	Signer crypto.Signer
	// CertValidationOpts contains various parameters for certificate chain
	// validation.
	CertValidationOpts CertValidationOpts
}

func New(origin string, signer crypto.Signer, cfg CertValidationConfig) (*ValidatedLogConfig, error) {
	vCfg := &ValidatedLogConfig{}

	if origin == "" {
		return nil, errors.New("empty origin")
	}
	vCfg.Origin = origin

	// Validate signer that only ECDSA is supported.
	// TODO(phboneff): if this is a library this should also allow RSA as per RFC6962.
	if signer == nil {
		return nil, errors.New("empty signer")
	}
	switch keyType := signer.Public().(type) {
	case *ecdsa.PublicKey:
	default:
		return nil, fmt.Errorf("unsupported key type: %v", keyType)
	}
	vCfg.Signer = signer

	vlc, err := newCertValidationOpts(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid cert validation config: %v", err)
	}
	vCfg.CertValidationOpts = *vlc

	return vCfg, nil
}

// newCertValidationOpts checks that a log validation config is valid,
// parses it and loads necessary resources.
func newCertValidationOpts(cfg CertValidationConfig) (*CertValidationOpts, error) {
	// Load the trusted roots.
	if len(cfg.RootsPemFile) == 0 {
		return nil, errors.New("empty rootsPemFile")
	}
	roots := x509util.NewPEMCertPool()
	if err := roots.AppendCertsFromPEMFile(cfg.RootsPemFile); err != nil {
		return nil, fmt.Errorf("failed to read trusted roots: %v", err)
	}

	if cfg.RejectExpired && cfg.RejectUnexpired {
		return nil, errors.New("rejecting all certificates")
	}

	// Validate the time interval.
	if cfg.NotAfterStart != nil && cfg.NotAfterLimit != nil && (cfg.NotAfterLimit).Before(*cfg.NotAfterStart) {
		return nil, errors.New("limit before start")
	}

	validationOpts := CertValidationOpts{
		trustedRoots:    roots,
		rejectExpired:   cfg.RejectExpired,
		rejectUnexpired: cfg.RejectUnexpired,
		notAfterStart:   cfg.NotAfterStart,
		notAfterLimit:   cfg.NotAfterLimit,
	}

	// Filter which extended key usages are allowed.
	lExtKeyUsages := []string{}
	if cfg.ExtKeyUsages != "" {
		lExtKeyUsages = strings.Split(cfg.ExtKeyUsages, ",")
	}
	// Validate the extended key usages list.
	for _, kuStr := range lExtKeyUsages {
		if ku, ok := stringToKeyUsage[kuStr]; ok {
			// If "Any" is specified, then we can ignore the entire list and
			// just disable EKU checking.
			if ku == x509.ExtKeyUsageAny {
				klog.Info("Found ExtKeyUsageAny, allowing all EKUs")
				validationOpts.extKeyUsages = nil
				break
			}
			validationOpts.extKeyUsages = append(validationOpts.extKeyUsages, ku)
		} else {
			return nil, fmt.Errorf("unknown extended key usage: %s", kuStr)
		}
	}
	// Filter which extensions are rejected.
	var err error
	lRejectExtensions := []string{}
	if cfg.RejectExtensions != "" {
		lRejectExtensions = strings.Split(cfg.RejectExtensions, ",")
		validationOpts.rejectExtIds, err = parseOIDs(lRejectExtensions)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RejectExtensions: %v", err)
		}
	}

	return &validationOpts, nil
}

func parseOIDs(oids []string) ([]asn1.ObjectIdentifier, error) {
	ret := make([]asn1.ObjectIdentifier, 0, len(oids))
	for _, s := range oids {
		bits := strings.Split(s, ".")
		var oid asn1.ObjectIdentifier
		for _, n := range bits {
			p, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			oid = append(oid, p)
		}
		ret = append(ret, oid)
	}
	return ret, nil
}

var stringToKeyUsage = map[string]x509.ExtKeyUsage{
	"Any":                        x509.ExtKeyUsageAny,
	"ServerAuth":                 x509.ExtKeyUsageServerAuth,
	"ClientAuth":                 x509.ExtKeyUsageClientAuth,
	"CodeSigning":                x509.ExtKeyUsageCodeSigning,
	"EmailProtection":            x509.ExtKeyUsageEmailProtection,
	"IPSECEndSystem":             x509.ExtKeyUsageIPSECEndSystem,
	"IPSECTunnel":                x509.ExtKeyUsageIPSECTunnel,
	"IPSECUser":                  x509.ExtKeyUsageIPSECUser,
	"TimeStamping":               x509.ExtKeyUsageTimeStamping,
	"OCSPSigning":                x509.ExtKeyUsageOCSPSigning,
	"MicrosoftServerGatedCrypto": x509.ExtKeyUsageMicrosoftServerGatedCrypto,
	"NetscapeServerGatedCrypto":  x509.ExtKeyUsageNetscapeServerGatedCrypto,
}
