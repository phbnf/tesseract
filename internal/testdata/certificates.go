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

package testdata

import _ "embed"

// This file holds test certificates. It contains six issuance chains.
// TODO(phboneff): clean this and make use of a single chain if possible.

// Issuance chain 1
// ================
// The next section holds a self signed root, a leaf cert and a precert
// issued from it directly. These certs were generated with ./gen.go.

// CACertPEM is a valid test CA certificate.
//
// $ openssl x509  -in internal/testdata/test_root_ca_cert.pem -noout -text
//
// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 1 (0x1)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	Validity
//	    Not Before: Dec  5 18:05:50 2024 GMT
//	    Not After : Dec  5 18:05:50 2034 GMT
//	Subject: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (384 bit)
//	        pub:
//	            04:c5:e8:0a:7e:fd:d5:3a:3e:74:67:2d:29:60:dd:
//	            15:ad:eb:47:38:49:44:c6:87:33:53:a1:79:55:09:
//	            27:16:df:fb:61:48:7f:0d:17:25:6c:e2:93:22:7f:
//	            fe:2f:4e:52:29:d2:35:f8:d7:d3:22:4e:aa:3a:2a:
//	            7c:10:a8:4a:cb:3d:2c:82:e0:d4:4c:70:ca:df:5a:
//	            83:c4:85:df:bd:d5:c4:51:68:15:3e:f8:5d:60:55:
//	            45:c3:4f:0e:ec:94:dd
//	        ASN1 OID: secp384r1
//	        NIST CURVE: P-384
//	X509v3 extensions:
//	    X509v3 Key Usage: critical
//	        Certificate Sign, CRL Sign
//	    X509v3 Basic Constraints: critical
//	        CA:TRUE
//	    X509v3 Subject Key Identifier:
//	        77:1D:7C:21:61:2D:C2:05:7D:AA:30:1E:6B:7F:8F:9B:DC:61:20:68
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	30:66:02:31:00:9d:aa:cd:cf:a4:c2:a7:ce:4e:7e:52:25:67:
//	b6:9a:aa:fe:17:b5:6c:2c:d1:c7:6a:03:72:12:fc:df:53:f5:
//	1b:70:56:3a:a4:85:15:a0:57:a4:0e:5c:97:ba:d4:83:75:02:
//	31:00:d3:93:fd:a0:8a:88:3d:44:48:6b:a2:fe:27:74:42:df:
//	94:44:93:d6:31:99:90:24:20:4c:41:b1:f6:3a:1a:9a:51:bd:
//	bf:93:88:29:ee:1c:f5:e7:ad:85:f0:8f:46:fa
//
//go:embed test_root_ca_cert.pem
var CACertPEM string

// PrecertPEMValid is a test certificate containing a valid CT precertificate
// extension.
//
// $ openssl x509  -in internal/testdata/test_leaf_pre_cert_signed_by_root.pem -noout -text
//
// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 200 (0xc8)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	Validity
//	    Not Before: Dec  5 18:05:50 2024 GMT
//	    Not After : Dec  5 18:05:50 2025 GMT
//	Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (384 bit)
//	        pub:
//	            04:5b:04:6c:07:ac:05:1b:06:d5:e3:3c:7f:69:17:
//	            5b:67:f0:a6:c6:9a:61:ed:cc:b1:51:50:30:2a:07:
//	            0c:1c:a7:9b:fe:4e:48:43:5c:eb:88:5e:ce:a3:59:
//	            a7:5d:6f:c6:82:53:f5:f0:3e:09:ab:9e:c5:e9:c5:
//	            ba:bc:2f:39:8c:a7:63:86:d4:52:05:44:83:6e:1f:
//	            54:1f:cd:19:cc:d4:ec:3f:d1:7e:06:95:14:4c:fd:
//	            8f:ed:ba:7f:93:d4:f4
//	        ASN1 OID: secp384r1
//	        NIST CURVE: P-384
//	X509v3 extensions:
//	    X509v3 Key Usage: critical
//	        Digital Signature, Key Encipherment
//	    X509v3 Extended Key Usage:
//	        TLS Web Server Authentication
//	    X509v3 Basic Constraints: critical
//	        CA:FALSE
//	    X509v3 Authority Key Identifier:
//	        77:1D:7C:21:61:2D:C2:05:7D:AA:30:1E:6B:7F:8F:9B:DC:61:20:68
//	    X509v3 Subject Alternative Name:
//	        DNS:test.transparency.dev
//	    CT Precertificate Poison: critical
//	        NULL
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	30:64:02:30:19:40:dc:a6:8a:36:e8:8c:f6:85:02:15:58:f1:
//	31:18:b2:7d:0e:00:bd:05:b7:a4:fb:eb:0c:f7:3b:36:80:fa:
//	c2:5b:3b:33:49:49:d9:20:48:bc:db:23:93:0d:1c:2d:02:30:
//	68:7d:1a:6f:91:9f:32:15:fb:eb:20:74:96:21:69:0c:6e:16:
//	46:c5:26:0e:60:8b:26:bc:f8:7d:e2:d4:16:8e:31:ce:36:a1:
//	45:e0:9c:f3:f3:c4:56:97:9f:f8:db:6f
//
//go:embed test_leaf_pre_cert_signed_by_root.pem
var PrecertPEMValid string

// TestCertPEM is a certificate issued by CACertPEM, no CT extensions.
//
// $ openssl x509  -in internal/testdata/test_leaf_cert_signed_by_root.pem -noout -text
//
// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 100 (0x64)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	Validity
//	    Not Before: Dec  5 18:05:50 2024 GMT
//	    Not After : Dec  5 18:05:50 2025 GMT
//	Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (384 bit)
//	        pub:
//	            04:5b:04:6c:07:ac:05:1b:06:d5:e3:3c:7f:69:17:
//	            5b:67:f0:a6:c6:9a:61:ed:cc:b1:51:50:30:2a:07:
//	            0c:1c:a7:9b:fe:4e:48:43:5c:eb:88:5e:ce:a3:59:
//	            a7:5d:6f:c6:82:53:f5:f0:3e:09:ab:9e:c5:e9:c5:
//	            ba:bc:2f:39:8c:a7:63:86:d4:52:05:44:83:6e:1f:
//	            54:1f:cd:19:cc:d4:ec:3f:d1:7e:06:95:14:4c:fd:
//	            8f:ed:ba:7f:93:d4:f4
//	        ASN1 OID: secp384r1
//	        NIST CURVE: P-384
//	X509v3 extensions:
//	    X509v3 Key Usage: critical
//	        Digital Signature, Key Encipherment
//	    X509v3 Extended Key Usage:
//	        TLS Web Server Authentication
//	    X509v3 Basic Constraints: critical
//	        CA:FALSE
//	    X509v3 Authority Key Identifier:
//	        77:1D:7C:21:61:2D:C2:05:7D:AA:30:1E:6B:7F:8F:9B:DC:61:20:68
//	    X509v3 Subject Alternative Name:
//	        DNS:test.transparency.dev
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	30:65:02:30:1b:56:30:9c:4c:bd:9e:18:13:2f:08:c0:52:d5:
//	83:29:8e:f9:22:77:d9:77:9e:be:20:2a:cc:5a:4b:46:ea:71:
//	b6:91:21:72:fd:db:c0:a1:9c:a5:69:a1:27:fa:1b:55:02:31:
//	00:d0:3f:2b:92:83:d9:a8:b6:d8:28:f3:7b:1e:4a:b8:fd:ec:
//	90:07:5c:02:9b:51:86:54:44:4b:cb:c9:b1:00:ac:93:ec:05:
//	7b:35:45:23:43:dd:47:dd:0e:d0:cf:98:f0
//
//go:embed test_leaf_cert_signed_by_root.pem
var TestCertPEM string

// Issuance chain 2
// ================
// The next section holds:
//   - an intermediate with the CT preissuer bit signed with the root above.
//   - a pre-cert issued by this intermediate

// PreIntermdiateFromRoot is an intermediate with the CT preissuer bit signed
// with the root above.
//
// $ openssl x509  -in internal/testdata/test_pre_intermediate_ca_cert.pem -noout -text
//
// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 2 (0x2)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	Validity
//	    Not Before: Dec  5 18:05:50 2024 GMT
//	    Not After : Dec  5 18:05:50 2029 GMT
//	Subject: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (384 bit)
//	        pub:
//	            04:3d:e3:59:7a:d0:8b:55:a7:96:30:88:6e:ba:2c:
//	            60:d5:30:38:b3:e9:da:62:19:c0:1f:b5:12:c2:fe:
//	            77:59:77:30:47:e3:e1:36:02:9b:5c:9c:7e:65:aa:
//	            56:76:91:02:0f:d8:64:aa:40:41:5e:19:fa:b2:39:
//	            de:13:a6:ee:1b:96:34:91:67:36:7e:7c:2f:cc:8e:
//	            c0:7f:e9:fb:b6:fa:d9:f9:1f:ed:3c:18:59:4d:a0:
//	            ab:ee:11:e3:f0:2c:87
//	        ASN1 OID: secp384r1
//	        NIST CURVE: P-384
//	X509v3 extensions:
//	    X509v3 Key Usage: critical
//	        Certificate Sign, CRL Sign
//	    X509v3 Basic Constraints: critical
//	        CA:TRUE
//	    X509v3 Subject Key Identifier:
//	        1F:FE:3D:85:AC:F5:38:C7:90:1C:6C:EA:E7:5F:45:74:83:CC:95:39
//	    X509v3 Authority Key Identifier:
//	        77:1D:7C:21:61:2D:C2:05:7D:AA:30:1E:6B:7F:8F:9B:DC:61:20:68
//	    CT Precertificate Signer:
//
// Signature Algorithm: ecdsa-with-SHA384
// Signature Value:
//
//	30:66:02:31:00:a1:62:a6:36:99:62:27:f4:e7:8b:9b:5e:ff:
//	80:4c:75:39:04:cc:80:d7:64:12:09:e8:80:e6:10:af:24:81:
//	2a:59:17:7a:58:da:6f:ca:f3:46:d3:5b:5c:e6:e1:dd:9c:02:
//	31:00:dd:c8:3a:b9:5d:9c:08:3c:27:73:11:17:fe:7a:82:98:
//	79:f8:a3:e3:16:54:c5:9f:79:d7:4f:1a:e2:55:48:d1:87:f2:
//	ab:2f:ad:81:dd:6e:b9:fc:59:77:37:6e:6e:75
//
//go:embed test_pre_intermediate_ca_cert.pem
var PreIntermediateFromRoot string

// PreCertFromPreIntermediate is a pre-cert issued by PreIntermediateFromRoot.
//
// $ openssl x509  -in internal/testdata/test_leaf_pre_cert_signed_by_pre_intermediate.pem -noout -text
//
// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 200 (0xc8)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//	Validity
//	    Not Before: Dec  5 18:05:50 2024 GMT
//	    Not After : Dec  5 18:05:50 2025 GMT
//	Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (384 bit)
//	        pub:
//	            04:46:10:60:6d:e5:70:0d:fa:8f:ea:8c:70:40:6e:
//	            eb:dd:15:88:8a:6e:94:54:ac:f7:92:77:53:68:65:
//	            c1:55:d4:c0:92:2e:b4:08:d9:07:50:d3:12:f4:fb:
//	            56:08:ff:38:32:41:35:6e:53:12:af:57:88:39:68:
//	            81:e0:1b:4c:82:4a:de:ac:52:d4:46:a7:a2:55:73:
//	            78:7a:fd:98:0f:bb:88:5b:bc:f6:7b:9a:77:49:11:
//	            ec:e6:1b:f3:c3:76:4a
//	        ASN1 OID: secp384r1
//	        NIST CURVE: P-384
//	X509v3 extensions:
//	    X509v3 Key Usage: critical
//	        Digital Signature, Key Encipherment
//	    X509v3 Extended Key Usage:
//	        TLS Web Server Authentication
//	    X509v3 Basic Constraints: critical
//	        CA:FALSE
//	    X509v3 Authority Key Identifier:
//	        1F:FE:3D:85:AC:F5:38:C7:90:1C:6C:EA:E7:5F:45:74:83:CC:95:39
//	    X509v3 Subject Alternative Name:
//	        DNS:test.transparency.dev
//	    CT Precertificate Poison: critical
//	        NULL
//
// Signature Algorithm: ecdsa-with-SHA384
// Signature Value:
//
//	30:66:02:31:00:af:dc:05:cf:bc:09:c1:d1:a4:26:3f:29:87:
//	87:ba:c9:e9:4c:d3:a6:06:c3:7c:64:0f:11:fe:d2:02:5c:50:
//	3c:bf:5a:9f:b7:8f:d9:df:44:0e:15:08:27:90:b3:8c:57:02:
//	31:00:ea:02:f2:78:e5:99:2d:9a:26:af:c5:49:da:42:9f:71:
//	63:12:db:6d:85:55:43:d2:26:66:fd:5d:81:71:13:50:2d:69:
//	cf:76:d7:05:3e:d6:04:3c:39:e7:20:7a:21:c2
//
//go:embed test_leaf_pre_cert_signed_by_pre_intermediate.pem
var PreCertFromPreIntermediate string

// CertFromPreIntermediate is a cert issued by PreIntermediateFromRoot.
// This is *NOT* a PreCert, as opposed to PreCertFromPreIntermediate.
// These certs should not be a thing in the real world, and we only use
// this one in tests.
//
// $ openssl x509  -in internal/testdata/test_leaf_cert_signed_by_pre_intermediate.pem -noout -text
//
// Data:
//
//   Version: 3 (0x2)
//   Serial Number: 100 (0x64)
//   Signature Algorithm: ecdsa-with-SHA384
//   Issuer: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//   Validity
//   	Not Before: Dec  5 18:05:50 2024 GMT
//   	Not After : Dec  5 18:05:50 2025 GMT
//   Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//   Subject Public Key Info:
//   	Public Key Algorithm: id-ecPublicKey
//   		Public-Key: (384 bit)
//   		pub:
//   			04:46:10:60:6d:e5:70:0d:fa:8f:ea:8c:70:40:6e:
//   			eb:dd:15:88:8a:6e:94:54:ac:f7:92:77:53:68:65:
//   			c1:55:d4:c0:92:2e:b4:08:d9:07:50:d3:12:f4:fb:
//   			56:08:ff:38:32:41:35:6e:53:12:af:57:88:39:68:
//   			81:e0:1b:4c:82:4a:de:ac:52:d4:46:a7:a2:55:73:
//   			78:7a:fd:98:0f:bb:88:5b:bc:f6:7b:9a:77:49:11:
//   			ec:e6:1b:f3:c3:76:4a
//   		ASN1 OID: secp384r1
//   		NIST CURVE: P-384
//   X509v3 extensions:
//   	X509v3 Key Usage: critical
//   		Digital Signature, Key Encipherment
//   	X509v3 Extended Key Usage:
//   		TLS Web Server Authentication
//   	X509v3 Basic Constraints: critical
//   		CA:FALSE
//   	X509v3 Authority Key Identifier:
//   		1F:FE:3D:85:AC:F5:38:C7:90:1C:6C:EA:E7:5F:45:74:83:CC:95:39
//   	X509v3 Subject Alternative Name:
//   		DNS:test.transparency.dev
//   Signature Algorithm: ecdsa-with-SHA384
//   Signature Value:
//   30:65:02:30:35:aa:8a:19:87:15:32:a8:a1:2f:a7:46:67:cb:
//   f5:ac:69:84:ef:2b:8d:e2:49:3f:a5:18:92:f6:e7:1d:b2:f2:
//   91:76:99:3f:ea:f3:b7:ec:df:94:96:78:b5:54:d2:59:02:31:
//   00:b0:11:91:9a:bb:6c:3b:f9:14:e4:1f:3e:b4:40:22:0b:46:
//   45:2a:18:cf:7c:20:b3:a7:56:3c:f6:c2:e0:7c:47:8a:a4:fd:
//   bf:6b:a2:48:c7:1e:4c:f8:f4:8a:df:5b:14

//go:embed test_leaf_cert_signed_by_pre_intermediate.pem
var CertFromPreIntermediate string

// Issuance chain 3
// ================
// The next section holds:
//   - an intermediate signed with the root above.
//   - a pre-cert issued by this intermediate
//   - a cert issued by this intermediate
//
// IntermediateFromRoot is an intermedidate cert signed by the root above.
//
// Certificate:
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number: 2 (0x2)
//	    Signature Algorithm: ecdsa-with-SHA384
//	    Issuer: C=GB, O=TrustFabric Transparency.dev Test Root Test CA, CN=TrustFabric Transparency.dev Test Root Test CA
//	    Validity
//	        Not Before: Dec  5 18:05:50 2024 GMT
//	        Not After : Dec  5 18:05:50 2029 GMT
//	    Subject: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//	    Subject Public Key Info:
//	        Public Key Algorithm: id-ecPublicKey
//	            Public-Key: (384 bit)
//	            pub:
//	                04:c6:3b:93:72:73:76:1b:f6:16:f6:93:2b:22:c0:
//	                1d:8e:e3:c9:1a:64:b8:42:58:0e:72:0a:38:49:84:
//	                d2:d3:e7:23:52:ee:9d:8a:28:65:73:2e:2e:7e:05:
//	                1c:d5:4f:5b:42:37:e3:bb:8f:54:3d:68:a9:c3:07:
//	                31:aa:4a:cc:8b:93:d6:a4:31:b2:8e:d0:aa:56:3b:
//	                49:ed:07:d3:36:ff:17:50:ad:6d:65:d7:5d:76:70:
//	                d5:08:f2:95:e0:11:0c
//	            ASN1 OID: secp384r1
//	            NIST CURVE: P-384
//	    X509v3 extensions:
//	        X509v3 Key Usage: critical
//	            Certificate Sign, CRL Sign
//	        X509v3 Basic Constraints: critical
//	            CA:TRUE
//	        X509v3 Subject Key Identifier:
//	            A0:D7:2B:CF:08:6F:C0:07:39:9B:C5:A9:87:1D:F7:CC:7D:6B:6F:29
//	        X509v3 Authority Key Identifier:
//	            77:1D:7C:21:61:2D:C2:05:7D:AA:30:1E:6B:7F:8F:9B:DC:61:20:68
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	    30:65:02:30:13:7b:99:45:f5:f5:c2:8b:bf:b4:83:8c:10:27:
//	    5e:50:a7:05:c0:61:8a:50:3f:76:2e:ec:88:71:d7:a7:a1:46:
//	    56:3b:3a:bc:e7:74:22:94:56:91:95:80:a5:a1:43:08:02:31:
//	    00:81:a0:12:84:45:6f:35:b3:3d:9b:98:ca:28:33:d2:b9:bf:
//	    8b:82:f7:a9:77:ee:2e:9f:90:0f:36:00:3e:c8:63:4c:1c:6c:
//	    de:e8:79:1a:32:44:4a:4e:47:6e:af:a3:24
//
//go:embed test_intermediate_ca_cert.pem
var IntermediateFromRoot string

// CertFromIntermediate is a leaf cert signed by the intermediate above.
//
// Certificate:
//
//	Data:
//	Version: 3 (0x2)
//	Serial Number: 100 (0x64)
//	Signature Algorithm: ecdsa-with-SHA384
//	Issuer: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//	Validity
//		Not Before: Dec  5 18:05:50 2024 GMT
//		Not After : Dec  5 18:05:50 2025 GMT
//	Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//	Subject Public Key Info:
//		Public Key Algorithm: id-ecPublicKey
//			Public-Key: (384 bit)
//			pub:
//				04:f4:15:a1:50:6c:d3:96:ad:9c:a0:f6:c0:90:4f:
//				05:13:64:2d:bf:2f:7a:86:4e:c8:25:c3:7d:9e:6f:
//				c3:44:b6:29:98:01:f4:d5:06:58:c9:cc:82:21:79:
//				97:88:3f:af:4c:bd:93:92:39:08:18:5f:81:c4:0b:
//				a0:ea:83:f8:6d:81:9a:68:20:bf:ad:2c:9b:1f:02:
//				08:cc:c2:16:a3:18:92:62:fa:b5:b0:da:ba:8b:98:
//				89:0a:d1:8c:65:3f:62
//			ASN1 OID: secp384r1
//			NIST CURVE: P-384
//	X509v3 extensions:
//		X509v3 Key Usage: critical
//			Digital Signature, Key Encipherment
//		X509v3 Extended Key Usage:
//			TLS Web Server Authentication
//		X509v3 Basic Constraints: critical
//			CA:FALSE
//		X509v3 Authority Key Identifier:
//			A0:D7:2B:CF:08:6F:C0:07:39:9B:C5:A9:87:1D:F7:CC:7D:6B:6F:29
//		X509v3 Subject Alternative Name:
//			DNS:test.transparency.dev
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	30:66:02:31:00:fd:08:f9:21:b5:a6:e0:32:aa:d0:aa:e2:07:
//	9c:fd:cc:26:b5:9a:bc:27:60:4f:ea:52:76:9f:cd:5c:23:b0:
//	fd:9e:5d:e9:73:a4:8a:1a:b5:b7:12:c2:69:e7:f1:bd:eb:02:
//	31:00:af:09:6b:61:78:6c:14:a3:9d:bd:e4:bf:91:43:a2:98:
//	a2:50:27:5d:2c:df:12:38:cd:b7:3d:d6:73:69:3a:5d:54:9c:
//	58:63:35:3c:39:78:26:37:08:75:3f:4b:fb:68
//
//go:embed test_leaf_cert_signed_by_intermediate.pem
var CertFromIntermediate string

// PreCertFromIntrmediate is a pre-cert signed by the intermediate above.
//
// Certificate:
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number: 200 (0xc8)
//	    Signature Algorithm: ecdsa-with-SHA384
//	    Issuer: C=GB, O=TrustFabric Transparency.dev Test Intermediate Test CA, CN=TrustFabric Transparency.dev Test Intermediate Test CA
//	    Validity
//	        Not Before: Dec  5 18:05:50 2024 GMT
//	        Not After : Dec  5 18:05:50 2025 GMT
//	    Subject: C=GB, ST=London, L=London, O=TrustFabric Transparency.dev Test, OU=TrustFabric, CN=test.transparency.dev
//	    Subject Public Key Info:
//	        Public Key Algorithm: id-ecPublicKey
//	            Public-Key: (384 bit)
//	            pub:
//	                04:f4:15:a1:50:6c:d3:96:ad:9c:a0:f6:c0:90:4f:
//	                05:13:64:2d:bf:2f:7a:86:4e:c8:25:c3:7d:9e:6f:
//	                c3:44:b6:29:98:01:f4:d5:06:58:c9:cc:82:21:79:
//	                97:88:3f:af:4c:bd:93:92:39:08:18:5f:81:c4:0b:
//	                a0:ea:83:f8:6d:81:9a:68:20:bf:ad:2c:9b:1f:02:
//	                08:cc:c2:16:a3:18:92:62:fa:b5:b0:da:ba:8b:98:
//	                89:0a:d1:8c:65:3f:62
//	            ASN1 OID: secp384r1
//	            NIST CURVE: P-384
//	    X509v3 extensions:
//	        X509v3 Key Usage: critical
//	            Digital Signature, Key Encipherment
//	        X509v3 Extended Key Usage:
//	            TLS Web Server Authentication
//	        X509v3 Basic Constraints: critical
//	            CA:FALSE
//	        X509v3 Authority Key Identifier:
//	            A0:D7:2B:CF:08:6F:C0:07:39:9B:C5:A9:87:1D:F7:CC:7D:6B:6F:29
//	        X509v3 Subject Alternative Name:
//	            DNS:test.transparency.dev
//	        CT Precertificate Poison: critical
//	            NULL
//	Signature Algorithm: ecdsa-with-SHA384
//	Signature Value:
//	    30:64:02:30:51:d4:2e:f7:e0:50:06:e5:a5:97:1c:d2:f9:4f:
//	    6e:c2:3b:e0:db:59:16:db:8d:1b:a8:c4:c6:b8:0a:4f:a3:0d:
//	    38:43:72:d7:f8:e6:60:e3:b8:44:f2:1f:37:56:30:cb:02:30:
//	    13:62:9c:60:c9:57:d1:b9:e0:43:f7:cf:2c:99:eb:04:84:f7:
//	    de:af:fd:d6:1a:63:90:14:4c:53:40:dd:28:0b:aa:69:59:87:
//	    78:8b:65:9e:00:63:75:7a:4c:a0:9f:ca
//
//go:embed test_leaf_pre_cert_signed_by_intermediate.pem
var PreCertFromIntermediate string

// Issuance chain 4
// ================
// The next section holds a self signed root, an intermediate, and a leaf cert.
//
// FakeCACertPEM is a test CA cert for testing.
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number:
//	        b6:31:d2:ac:21:ab:65:20
//	Signature Algorithm: sha256WithRSAEncryption
//	    Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	    Validity
//	        Not Before: Jul 11 12:23:26 2016 GMT
//	        Not After : Jul 11 12:23:26 2017 GMT
//	    Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	    Subject Public Key Info:
//	        Public Key Algorithm: rsaEncryption
//	            Public-Key: (2048 bit)
//	            Modulus:
//	                00:a5:41:9a:7a:2d:98:a3:b5:78:6f:15:21:db:0c:
//	                c1:0e:a1:f8:26:f5:b3:b2:67:85:dc:a1:e6:b7:83:
//	                6d:da:63:da:d0:f6:a3:ff:bc:43:f5:2b:9f:00:19:
//	                6e:6b:60:4b:43:20:6e:e2:cb:2e:b6:65:ed:9b:dc:
//	                80:c3:e1:5a:96:af:60:78:0e:0e:fb:8f:ea:3e:3d:
//	                c9:67:8f:a4:57:1c:ba:e4:f3:37:a9:2f:dd:11:9d:
//	                10:5d:e5:d6:ef:d4:3b:06:d9:34:43:42:bb:bb:be:
//	                43:40:2b:e3:b6:d1:b5:6c:58:12:34:96:14:d4:fc:
//	                49:79:c5:26:8c:24:7d:b3:12:f5:f6:3e:b7:41:46:
//	                6b:6d:3a:41:fd:7c:e3:b5:fc:96:6c:c6:cc:ad:8d:
//	                48:09:73:44:64:ea:4f:17:1d:0a:4b:14:5a:19:07:
//	                4a:32:0f:41:2e:e4:85:bd:a1:e1:9b:de:63:7c:3b:
//	                bc:ec:aa:93:2a:0b:a8:c7:24:34:54:42:38:a5:d1:
//	                0c:c4:f9:9e:7c:69:42:71:77:d7:95:aa:bb:13:3d:
//	                f3:cc:c7:5d:b3:fd:76:25:25:e3:da:14:0e:59:81:
//	                e8:2c:58:e8:09:29:7d:22:02:91:95:81:eb:55:6f:
//	                2f:17:b9:af:4a:f3:84:8b:24:6e:ea:14:6b:bb:90:
//	                84:35
//	            Exponent: 65537 (0x10001)
//	    X509v3 extensions:
//	        X509v3 Subject Key Identifier:
//	            01:02:03:04
//	        X509v3 Authority Key Identifier:
//	            keyid:01:02:03:04
//
//	        X509v3 Basic Constraints: critical
//	            CA:TRUE, pathlen:10
//	        X509v3 Key Usage: critical
//	            Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//	Signature Algorithm: sha256WithRSAEncryption
//	     92:be:33:eb:d5:d4:32:e7:9e:4e:65:2a:e8:3f:67:b8:f4:d7:
//	     34:ab:95:11:6a:5d:ba:fd:57:9b:94:6e:8d:20:be:fb:7a:e1:
//	     49:ca:39:ea:92:d3:81:5a:b1:87:a3:9f:50:a4:e0:1e:11:de:
//	     c4:d1:07:a1:ca:d1:97:1a:92:bd:73:9a:11:ec:6a:9a:52:11:
//	     2d:40:e1:3b:4f:3c:1f:81:3f:4c:ab:6a:02:84:4f:8b:18:36:
//	     7a:cc:5c:a9:0e:25:2b:cd:57:53:88:d9:eb:82:b1:ce:62:76:
//	     56:d4:23:9e:01:b3:6d:2b:49:ea:d4:3a:c2:f5:76:a7:b3:2d:
//	     24:97:6f:b4:1c:74:6b:95:85:f6:b5:41:56:82:3c:ed:be:96:
//	     1e:5e:6a:2d:7b:f7:fd:7d:6e:3f:fb:c2:ec:61:b3:7c:7f:3b:
//	     f5:9c:64:61:5f:02:93:87:cd:81:f9:7e:53:3e:c1:f5:79:85:
//	     f4:41:87:c7:ca:bd:af:ab:2b:a4:aa:a8:1d:2c:50:ad:23:8f:
//	     db:13:1d:71:8a:85:bd:ac:59:6c:c4:53:c5:71:0c:90:91:f3:
//	     0b:41:ef:da:6e:27:bb:09:57:9c:97:b9:d7:fc:20:96:c5:75:
//	     96:ce:2e:6c:a8:b6:6e:b0:4d:0f:3e:01:95:ea:8b:cd:ae:47:
//	     d0:d9:01:b7
const FakeCACertPEM = `
-----BEGIN CERTIFICATE-----
MIIDrDCCApSgAwIBAgIJALYx0qwhq2UgMA0GCSqGSIb3DQEBCwUAMHExCzAJBgNV
BAYTAkdCMQ8wDQYDVQQIDAZMb25kb24xDzANBgNVBAcMBkxvbmRvbjEPMA0GA1UE
CgwGR29vZ2xlMQwwCgYDVQQLDANFbmcxITAfBgNVBAMMGEZha2VDZXJ0aWZpY2F0
ZUF1dGhvcml0eTAeFw0xNjA3MTExMjIzMjZaFw0xNzA3MTExMjIzMjZaMHExCzAJ
BgNVBAYTAkdCMQ8wDQYDVQQIDAZMb25kb24xDzANBgNVBAcMBkxvbmRvbjEPMA0G
A1UECgwGR29vZ2xlMQwwCgYDVQQLDANFbmcxITAfBgNVBAMMGEZha2VDZXJ0aWZp
Y2F0ZUF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKVB
mnotmKO1eG8VIdsMwQ6h+Cb1s7Jnhdyh5reDbdpj2tD2o/+8Q/UrnwAZbmtgS0Mg
buLLLrZl7ZvcgMPhWpavYHgODvuP6j49yWePpFccuuTzN6kv3RGdEF3l1u/UOwbZ
NENCu7u+Q0Ar47bRtWxYEjSWFNT8SXnFJowkfbMS9fY+t0FGa206Qf1847X8lmzG
zK2NSAlzRGTqTxcdCksUWhkHSjIPQS7khb2h4ZveY3w7vOyqkyoLqMckNFRCOKXR
DMT5nnxpQnF315WquxM988zHXbP9diUl49oUDlmB6CxY6AkpfSICkZWB61VvLxe5
r0rzhIskbuoUa7uQhDUCAwEAAaNHMEUwDQYDVR0OBAYEBAECAwQwDwYDVR0jBAgw
BoAEAQIDBDASBgNVHRMBAf8ECDAGAQH/AgEKMA8GA1UdDwEB/wQFAwMH/4AwDQYJ
KoZIhvcNAQELBQADggEBAJK+M+vV1DLnnk5lKug/Z7j01zSrlRFqXbr9V5uUbo0g
vvt64UnKOeqS04FasYejn1Ck4B4R3sTRB6HK0Zcakr1zmhHsappSES1A4TtPPB+B
P0yragKET4sYNnrMXKkOJSvNV1OI2euCsc5idlbUI54Bs20rSerUOsL1dqezLSSX
b7QcdGuVhfa1QVaCPO2+lh5eai179/19bj/7wuxhs3x/O/WcZGFfApOHzYH5flM+
wfV5hfRBh8fKva+rK6SqqB0sUK0jj9sTHXGKhb2sWWzEU8VxDJCR8wtB79puJ7sJ
V5yXudf8IJbFdZbOLmyotm6wTQ8+AZXqi82uR9DZAbc=
-----END CERTIFICATE-----`

// FakeIntermediateCertPEM is a test intermediate CA cert.
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number: 4792439526061490155 (0x42822a5b866fbfeb)
//	Signature Algorithm: sha256WithRSAEncryption
//	    Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	    Validity
//	        Not Before: May 13 14:26:44 2016 GMT
//	        Not After : Jul 12 14:26:44 2019 GMT
//	    Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//	    Subject Public Key Info:
//	        Public Key Algorithm: rsaEncryption
//	            Public-Key: (2048 bit)
//	            Modulus:
//	                00:ca:a4:0c:7a:6d:e9:26:22:d4:67:19:c8:29:40:
//	                c6:bd:cb:44:39:e7:fa:84:01:1d:b3:04:15:48:37:
//	                fa:55:d5:98:4b:2a:ff:14:0e:d6:ce:27:6b:29:d5:
//	                e8:8d:39:eb:be:97:be:53:21:d2:a3:f2:27:ef:46:
//	                68:1c:6f:84:77:85:b4:68:78:7a:d4:3d:50:49:89:
//	                8f:9e:6b:4a:ce:74:c0:0f:c8:68:38:7e:ae:82:ae:
//	                91:0c:6d:87:24:c4:48:f3:e0:8e:a8:3e:0c:f8:e1:
//	                e8:7f:a1:dd:29:f4:d0:eb:3a:b2:38:77:0f:1a:4e:
//	                a6:14:c4:b1:db:5b:ed:f9:a4:f0:9d:1e:d8:a8:d0:
//	                40:28:d6:fc:69:44:0b:37:37:e7:d6:fd:29:b0:70:
//	                36:47:00:89:81:5a:c9:51:cf:2d:a0:80:76:fc:d8:
//	                57:28:87:81:71:e4:10:4b:39:16:51:f2:85:ed:a0:
//	                34:41:bf:f3:52:28:f1:cd:c4:dc:31:f9:26:14:fd:
//	                b6:65:51:2f:76:e9:82:94:fc:2a:be:1a:a0:58:54:
//	                d8:b5:de:e3:96:08:07:50:3d:0e:35:26:e5:3a:c7:
//	                67:e8:8d:b6:f1:34:61:f6:0c:47:d2:fd:0b:51:cf:
//	                a6:99:97:d4:26:a1:12:14:dd:a2:0e:e5:68:4d:75:
//	                f7:c5
//	            Exponent: 65537 (0x10001)
//	    X509v3 extensions:
//	        X509v3 Authority Key Identifier:
//	            keyid:01:02:03:04
//
//	        X509v3 Basic Constraints: critical
//	            CA:TRUE, pathlen:0
//	        X509v3 Key Usage: critical
//	            Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//	Signature Algorithm: sha256WithRSAEncryption
//	     01:e2:3a:0c:00:bc:4c:e1:ac:d3:10:54:0c:fc:6b:e4:ac:c8:
//	     c2:00:05:74:39:3f:c5:9b:25:e1:e3:90:88:a9:13:8f:b9:66:
//	     99:2b:65:55:ea:f6:9f:30:39:d9:18:9c:e1:f1:e1:63:62:f4:
//	     f5:46:41:b2:c6:f4:8b:9f:87:d7:e9:93:c7:32:c9:15:83:8b:
//	     e5:76:d3:f0:8d:36:d6:b0:32:ad:c2:95:5d:dd:58:2f:7c:4e:
//	     3e:16:5f:f0:57:0c:27:98:da:32:b8:8d:81:95:f9:db:38:dc:
//	     76:15:d1:3a:01:9a:fb:eb:71:ca:bf:53:bc:d8:30:61:5c:42:
//	     22:81:0a:5c:f9:6d:31:3e:18:cb:eb:65:67:0e:e4:0f:cb:87:
//	     7f:22:d9:84:85:d6:2f:12:7c:35:67:00:e0:65:02:06:66:96:
//	     57:21:78:7a:46:b1:67:d2:9d:db:88:96:55:2f:4e:c4:6f:10:
//	     8b:1a:6a:a7:d5:2e:5e:50:a5:15:c1:3a:af:2d:6e:32:bc:e7:
//	     fd:a0:e9:e6:ab:d6:8c:4f:84:9d:70:f6:17:6c:f9:64:c5:5e:
//	     49:87:91:6b:ca:25:e6:d8:d7:7b:77:39:f4:a3:03:28:5a:45:
//	     2b:7c:85:dc:c3:cc:74:c5:c2:33:e3:1d:3f:21:e9:d5:3b:fe:
//	     13:1d:91:48
const FakeIntermediateCertPEM = `
-----BEGIN CERTIFICATE-----
MIIDnTCCAoWgAwIBAgIIQoIqW4Zvv+swDQYJKoZIhvcNAQELBQAwcTELMAkGA1UE
BhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYDVQQK
DAZHb29nbGUxDDAKBgNVBAsMA0VuZzEhMB8GA1UEAwwYRmFrZUNlcnRpZmljYXRl
QXV0aG9yaXR5MB4XDTE2MDUxMzE0MjY0NFoXDTE5MDcxMjE0MjY0NFowcjELMAkG
A1UEBhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYD
VQQKDAZHb29nbGUxDDAKBgNVBAsMA0VuZzEiMCAGA1UEAwwZRmFrZUludGVybWVk
aWF0ZUF1dGhvcml0eTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMqk
DHpt6SYi1GcZyClAxr3LRDnn+oQBHbMEFUg3+lXVmEsq/xQO1s4naynV6I05676X
vlMh0qPyJ+9GaBxvhHeFtGh4etQ9UEmJj55rSs50wA/IaDh+roKukQxthyTESPPg
jqg+DPjh6H+h3Sn00Os6sjh3DxpOphTEsdtb7fmk8J0e2KjQQCjW/GlECzc359b9
KbBwNkcAiYFayVHPLaCAdvzYVyiHgXHkEEs5FlHyhe2gNEG/81Io8c3E3DH5JhT9
tmVRL3bpgpT8Kr4aoFhU2LXe45YIB1A9DjUm5TrHZ+iNtvE0YfYMR9L9C1HPppmX
1CahEhTdog7laE1198UCAwEAAaM4MDYwDwYDVR0jBAgwBoAEAQIDBDASBgNVHRMB
Af8ECDAGAQH/AgEAMA8GA1UdDwEB/wQFAwMH/4AwDQYJKoZIhvcNAQELBQADggEB
AAHiOgwAvEzhrNMQVAz8a+SsyMIABXQ5P8WbJeHjkIipE4+5ZpkrZVXq9p8wOdkY
nOHx4WNi9PVGQbLG9Iufh9fpk8cyyRWDi+V20/CNNtawMq3ClV3dWC98Tj4WX/BX
DCeY2jK4jYGV+ds43HYV0ToBmvvrccq/U7zYMGFcQiKBClz5bTE+GMvrZWcO5A/L
h38i2YSF1i8SfDVnAOBlAgZmllcheHpGsWfSnduIllUvTsRvEIsaaqfVLl5QpRXB
Oq8tbjK85/2g6ear1oxPhJ1w9hds+WTFXkmHkWvKJebY13t3OfSjAyhaRSt8hdzD
zHTFwjPjHT8h6dU7/hMdkUg=
-----END CERTIFICATE-----`

// LeafSignedByFakeIntermediateCertPEM is a test cert signed by the intermediate CA.
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number: 4792439526061490155 (0x42822a5b866fbfeb)
//	Signature Algorithm: sha256WithRSAEncryption
//	    Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//	    Validity
//	        Not Before: May 13 14:26:44 2016 GMT
//	        Not After : Jul 12 14:26:44 2019 GMT
//	    Subject: C=US, ST=California, L=Mountain View, O=Google Inc, CN=*.google.com, SN=RFC5280 s4.2.1.9 'The pathLenConstraint field ... gives the maximum number of non-self-issued intermediate certificates that may follow this certificate in a valid certification path.', GN=Intermediate CA cert used to sign
//	    Subject Public Key Info:
//	        Public Key Algorithm: id-ecPublicKey
//	            Public-Key: (256 bit)
//	                04:c4:09:39:84:f5:15:8d:12:54:b2:02:9c:f9:01:
//	                e2:6d:35:47:d4:0d:d0:11:61:66:09:35:1d:cb:12:
//	                14:95:b2:3f:ff:35:bd:22:8e:4d:fc:38:50:2d:22:
//	                d6:98:1e:ca:a0:23:af:a4:96:7e:32:d1:82:5f:31:
//	                57:fb:28:ff:37
//	            ASN1 OID: prime256v1
//	            NIST CURVE: P-256
//	    X509v3 extensions:
//	        X509v3 Extended Key Usage:
//	            TLS Web Server Authentication, TLS Web Client Authentication
//	        X509v3 Subject Alternative Name:
//	            DNS:*.google.com, DNS:*.android.com, DNS:*.appengine.google.com, DNS:*.cloud.google.com,
//	            DNS:*.google-analytics.com, DNS:*.google.ca, DNS:*.google.cl, DNS:*.google.co.in, DNS:*.google.co.jp,
//	            DNS:*.google.co.uk, DNS:*.google.com.ar, DNS:*.google.com.au, DNS:*.google.com.br, DNS:*.google.com.co,
//	            DNS:*.google.com.mx, DNS:*.google.com.tr, DNS:*.google.com.vn, DNS:*.google.de, DNS:*.google.es,
//	            DNS:*.google.fr, DNS:*.google.hu, DNS:*.google.it, DNS:*.google.nl, DNS:*.google.pl, DNS:*.google.pt,
//	            DNS:*.googleadapis.com, DNS:*.googleapis.cn, DNS:*.googlecommerce.com, DNS:*.googlevideo.com,
//	            DNS:*.gstatic.cn, DNS:*.gstatic.com, DNS:*.gvt1.com, DNS:*.gvt2.com, DNS:*.metric.gstatic.com,
//	            DNS:*.urchin.com, DNS:*.url.google.com, DNS:*.youtube-nocookie.com, DNS:*.youtube.com,
//	            DNS:*.youtubeeducation.com, DNS:*.ytimg.com, DNS:android.clients.google.com, DNS:android.com, DNS:g.co,
//	            DNS:goo.gl, DNS:google-analytics.com, DNS:google.com, DNS:googlecommerce.com, DNS:urchin.com,
//	            DNS:youtu.be, DNS:youtube.com, DNS:youtubeeducation.com
//	        X509v3 Key Usage:
//	            Digital Signature
//	        Authority Information Access:
//	            CA Issuers - URI:http://pki.google.com/GIAG2.crt
//	            OCSP - URI:http://clients1.google.com/ocsp
//
//	        X509v3 Subject Key Identifier:
//	            DB:F4:6E:63:EE:E2:DC:BE:BF:38:60:4F:98:31:D0:64:44:F1:63:D8
//	        X509v3 Basic Constraints: critical
//	            CA:FALSE
//	        X509v3 Certificate Policies:
//	            Policy: 1.3.6.1.4.1.11129.2.5.1
//	            Policy: 2.23.140.1.2.2
//
//	        X509v3 CRL Distribution Points:
//
//	            Full Name:
//	              URI:http://pki.google.com/GIAG2.crl
//
//	Signature Algorithm: sha256WithRSAEncryption
//	     0e:a6:6f:79:7d:38:4b:60:f0:c1:76:9c:4e:92:f5:24:ce:12:
//	     34:72:94:95:8d:cf:1c:0c:d6:78:6b:ee:66:2b:50:36:22:7a:
//	     be:ff:22:c7:dd:93:2c:40:83:2f:a0:37:29:8f:bb:98:22:bf:
//	     8e:c6:6c:b4:8b:8f:e9:1e:0f:bd:8a:df:df:f5:c9:aa:79:ac:
//	     00:e6:ca:a6:1a:74:8e:67:f9:5f:09:82:3c:f9:b4:5b:30:85:
//	     0b:ae:28:c2:b8:9c:23:7c:6a:59:66:ca:8e:bd:20:6e:20:e4:
//	     b3:46:f8:06:56:99:5c:b3:47:62:b6:e4:f6:92:10:85:ae:46:
//	     e5:c1:af:c1:a8:8a:b3:b6:f3:fb:2e:e1:26:56:98:e4:aa:de:
//	     29:0b:71:ef:0f:45:d4:c6:ce:4f:21:d6:59:18:89:df:7a:ac:
//	     a6:93:97:de:45:e5:87:06:e3:c7:a4:f2:14:39:b2:b1:99:0b:
//	     7e:85:cc:3a:62:c1:c4:fb:40:7c:e1:7b:71:f4:13:1e:e2:aa:
//	     94:7e:ba:a6:b5:65:e7:f6:e9:c1:c3:1a:92:62:c0:aa:c4:74:
//	     29:43:ee:f4:a6:6b:81:c6:50:7d:b3:a2:d2:b4:8c:c4:f6:cc:
//	     9a:0e:65:32:8f:14:65:8c:a0:30:20:d5:7a:cf:48:fb:84:a4:
//	     3a:30:fa:44
const LeafSignedByFakeIntermediateCertPEM = `
-----BEGIN CERTIFICATE-----
MIIH6DCCBtCgAwIBAgIIQoIqW4Zvv+swDQYJKoZIhvcNAQELBQAwcjELMAkGA1UE
BhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9uZG9uMQ8wDQYDVQQK
DAZHb29nbGUxDDAKBgNVBAsMA0VuZzEiMCAGA1UEAwwZRmFrZUludGVybWVkaWF0
ZUF1dGhvcml0eTAeFw0xNjA1MTMxNDI2NDRaFw0xOTA3MTIxNDI2NDRaMIIBWDEL
MAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDU1vdW50
YWluIFZpZXcxEzARBgNVBAoMCkdvb2dsZSBJbmMxFTATBgNVBAMMDCouZ29vZ2xl
LmNvbTGBwzCBwAYDVQQEDIG4UkZDNTI4MCBzNC4yLjEuOSAnVGhlIHBhdGhMZW5D
b25zdHJhaW50IGZpZWxkIC4uLiBnaXZlcyB0aGUgbWF4aW11bSBudW1iZXIgb2Yg
bm9uLXNlbGYtaXNzdWVkIGludGVybWVkaWF0ZSBjZXJ0aWZpY2F0ZXMgdGhhdCBt
YXkgZm9sbG93IHRoaXMgY2VydGlmaWNhdGUgaW4gYSB2YWxpZCBjZXJ0aWZpY2F0
aW9uIHBhdGguJzEqMCgGA1UEKgwhSW50ZXJtZWRpYXRlIENBIGNlcnQgdXNlZCB0
byBzaWduMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAExAk5hPUVjRJUsgKc+QHi
bTVH1A3QEWFmCTUdyxIUlbI//zW9Io5N/DhQLSLWmB7KoCOvpJZ+MtGCXzFX+yj/
N6OCBGMwggRfMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjCCA0IGA1Ud
EQSCAzkwggM1ggwqLmdvb2dsZS5jb22CDSouYW5kcm9pZC5jb22CFiouYXBwZW5n
aW5lLmdvb2dsZS5jb22CEiouY2xvdWQuZ29vZ2xlLmNvbYIWKi5nb29nbGUtYW5h
bHl0aWNzLmNvbYILKi5nb29nbGUuY2GCCyouZ29vZ2xlLmNsgg4qLmdvb2dsZS5j
by5pboIOKi5nb29nbGUuY28uanCCDiouZ29vZ2xlLmNvLnVrgg8qLmdvb2dsZS5j
b20uYXKCDyouZ29vZ2xlLmNvbS5hdYIPKi5nb29nbGUuY29tLmJygg8qLmdvb2ds
ZS5jb20uY2+CDyouZ29vZ2xlLmNvbS5teIIPKi5nb29nbGUuY29tLnRygg8qLmdv
b2dsZS5jb20udm6CCyouZ29vZ2xlLmRlggsqLmdvb2dsZS5lc4ILKi5nb29nbGUu
ZnKCCyouZ29vZ2xlLmh1ggsqLmdvb2dsZS5pdIILKi5nb29nbGUubmyCCyouZ29v
Z2xlLnBsggsqLmdvb2dsZS5wdIISKi5nb29nbGVhZGFwaXMuY29tgg8qLmdvb2ds
ZWFwaXMuY26CFCouZ29vZ2xlY29tbWVyY2UuY29tghEqLmdvb2dsZXZpZGVvLmNv
bYIMKi5nc3RhdGljLmNugg0qLmdzdGF0aWMuY29tggoqLmd2dDEuY29tggoqLmd2
dDIuY29tghQqLm1ldHJpYy5nc3RhdGljLmNvbYIMKi51cmNoaW4uY29tghAqLnVy
bC5nb29nbGUuY29tghYqLnlvdXR1YmUtbm9jb29raWUuY29tgg0qLnlvdXR1YmUu
Y29tghYqLnlvdXR1YmVlZHVjYXRpb24uY29tggsqLnl0aW1nLmNvbYIaYW5kcm9p
ZC5jbGllbnRzLmdvb2dsZS5jb22CC2FuZHJvaWQuY29tggRnLmNvggZnb28uZ2yC
FGdvb2dsZS1hbmFseXRpY3MuY29tggpnb29nbGUuY29tghJnb29nbGVjb21tZXJj
ZS5jb22CCnVyY2hpbi5jb22CCHlvdXR1LmJlggt5b3V0dWJlLmNvbYIUeW91dHVi
ZWVkdWNhdGlvbi5jb20wDAYDVR0PBAUDAweAADBoBggrBgEFBQcBAQRcMFowKwYI
KwYBBQUHMAKGH2h0dHA6Ly9wa2kuZ29vZ2xlLmNvbS9HSUFHMi5jcnQwKwYIKwYB
BQUHMAGGH2h0dHA6Ly9jbGllbnRzMS5nb29nbGUuY29tL29jc3AwHQYDVR0OBBYE
FNv0bmPu4ty+vzhgT5gx0GRE8WPYMAwGA1UdEwEB/wQCMAAwIQYDVR0gBBowGDAM
BgorBgEEAdZ5AgUBMAgGBmeBDAECAjAwBgNVHR8EKTAnMCWgI6Ahhh9odHRwOi8v
cGtpLmdvb2dsZS5jb20vR0lBRzIuY3JsMA0GCSqGSIb3DQEBCwUAA4IBAQAOpm95
fThLYPDBdpxOkvUkzhI0cpSVjc8cDNZ4a+5mK1A2Inq+/yLH3ZMsQIMvoDcpj7uY
Ir+Oxmy0i4/pHg+9it/f9cmqeawA5sqmGnSOZ/lfCYI8+bRbMIULrijCuJwjfGpZ
ZsqOvSBuIOSzRvgGVplcs0dituT2khCFrkblwa/BqIqztvP7LuEmVpjkqt4pC3Hv
D0XUxs5PIdZZGInfeqymk5feReWHBuPHpPIUObKxmQt+hcw6YsHE+0B84Xtx9BMe
4qqUfrqmtWXn9unBwxqSYsCqxHQpQ+70pmuBxlB9s6LStIzE9syaDmUyjxRljKAw
INV6z0j7hKQ6MPpE
-----END CERTIFICATE-----`

// Issuance chain 5
// ================
// The next section holds a self signed root, intermediate certs
// with various policy constraints, and a leaf cert.

// The matching certificate and private key are stored in pem format
// in this directory.
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number: 67554046 (0x406cafe)
//	Signature Algorithm: ecdsa-with-SHA256
//	    Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	    Validity
//	        Not Before: Dec  7 15:13:36 2016 GMT
//	        Not After : Dec  5 15:13:36 2026 GMT
//	    Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	    Subject Public Key Info:
//	        Public Key Algorithm: id-ecPublicKey
//	            Public-Key: (256 bit)
//	            pub:
//	                04:f2:d3:07:ef:7e:df:cf:ce:f4:f4:0a:5b:bc:9e:
//	                3f:cb:1c:fd:0c:46:dc:85:fb:c1:f6:d3:b2:ba:1d:
//	                51:f1:98:6c:48:a8:15:46:45:63:ca:df:d6:c9:ac:
//	                cf:60:3b:c7:4e:dd:b8:d2:16:ab:a0:09:24:1d:09:
//	                66:1e:4d:eb:a1
//	            ASN1 OID: prime256v1
//	            NIST CURVE: P-256
//	    X509v3 extensions:
//	        X509v3 Subject Key Identifier:
//	            01:02:03:04
//	        X509v3 Authority Key Identifier:
//	            keyid:01:02:03:04
//	        X509v3 Basic Constraints: critical
//	            CA:TRUE, pathlen:10
//	        X509v3 Key Usage: critical
//	            Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//	Signature Algorithm: ecdsa-with-SHA256
//	     30:46:02:21:00:a6:28:49:39:43:6f:80:e4:43:a6:1e:3b:aa:
//	     89:5e:c2:25:60:2a:e1:39:bd:55:43:ae:4d:5c:a9:a6:ef:ac:
//	     65:02:21:00:c9:c5:08:c6:59:93:b4:86:70:a5:6b:54:2b:5b:
//	     fc:0c:88:6b:b0:23:07:2b:c7:0c:27:de:87:2d:96:80:d5:56
const FakeRootCACertPEM = `
-----BEGIN CERTIFICATE-----
MIICHDCCAcGgAwIBAgIEBAbK/jAKBggqhkjOPQQDAjBxMQswCQYDVQQGEwJHQjEP
MA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdvb2ds
ZTEMMAoGA1UECxMDRW5nMSEwHwYDVQQDExhGYWtlQ2VydGlmaWNhdGVBdXRob3Jp
dHkwHhcNMTYxMjA3MTUxMzM2WhcNMjYxMjA1MTUxMzM2WjBxMQswCQYDVQQGEwJH
QjEPMA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdv
b2dsZTEMMAoGA1UECxMDRW5nMSEwHwYDVQQDExhGYWtlQ2VydGlmaWNhdGVBdXRo
b3JpdHkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATy0wfvft/PzvT0Clu8nj/L
HP0MRtyF+8H207K6HVHxmGxIqBVGRWPK39bJrM9gO8dO3bjSFqugCSQdCWYeTeuh
o0cwRTANBgNVHQ4EBgQEAQIDBDAPBgNVHSMECDAGgAQBAgMEMBIGA1UdEwEB/wQI
MAYBAf8CAQowDwYDVR0PAQH/BAUDAwf/gDAKBggqhkjOPQQDAgNJADBGAiEApihJ
OUNvgORDph47qolewiVgKuE5vVVDrk1cqabvrGUCIQDJxQjGWZO0hnCla1QrW/wM
iGuwIwcrxwwn3octloDVVg==
-----END CERTIFICATE-----`

// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 1111638594 (0x42424242)
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	Validity
//	    Not Before: Feb 13 09:33:59 2018 GMT
//	    Not After : Dec 23 09:33:59 2027 GMT
//	Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (256 bit)
//	        pub:
//	            04:f1:bf:2d:e8:8c:66:40:e3:a8:d1:54:e0:42:49:
//	            02:cb:dd:47:08:85:c2:67:41:4c:eb:f7:87:cd:8d:
//	            a3:09:c8:18:cc:2e:30:53:16:32:aa:d5:9c:08:73:
//	            c6:76:fa:fa:3a:38:e9:34:35:9c:51:d1:ee:12:81:
//	            5d:98:5f:5d:5d
//	        ASN1 OID: prime256v1
//	        NIST CURVE: P-256
//	X509v3 extensions:
//	    X509v3 Subject Key Identifier:
//	        01:02:03:04
//	    X509v3 Authority Key Identifier:
//	        keyid:01:02:03:04
//	    X509v3 Basic Constraints: critical
//	        CA:TRUE, pathlen:10
//	    X509v3 Policy Constraints: critical
//	        Require Explicit Policy:0
//	    X509v3 Key Usage: critical
//	        Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	30:44:02:20:4c:aa:27:8f:d9:83:32:76:40:17:a1:a8:00:1d:
//	bc:d1:45:b2:53:c6:47:77:48:f1:c3:89:68:5d:f4:7f:5c:52:
//	02:20:39:68:40:5c:fd:f0:2a:e2:3f:34:45:b3:19:2d:e3:4d:
//	58:cd:76:42:19:09:cf:5c:1c:e5:f1:71:e0:39:62:b9
const FakeIntermediateWithPolicyConstraintsCertPEM = `
-----BEGIN CERTIFICATE-----
MIICLDCCAdOgAwIBAgIEQkJCQjAKBggqhkjOPQQDAjBxMQswCQYDVQQGEwJHQjEP
MA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdvb2ds
ZTEMMAoGA1UECxMDRW5nMSEwHwYDVQQDExhGYWtlQ2VydGlmaWNhdGVBdXRob3Jp
dHkwHhcNMTgwMjEzMDkzMzU5WhcNMjcxMjIzMDkzMzU5WjByMQswCQYDVQQGEwJH
QjEPMA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdv
b2dsZTEMMAoGA1UECxMDRW5nMSIwIAYDVQQDExlGYWtlSW50ZXJtZWRpYXRlQXV0
aG9yaXR5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8b8t6IxmQOOo0VTgQkkC
y91HCIXCZ0FM6/eHzY2jCcgYzC4wUxYyqtWcCHPGdvr6OjjpNDWcUdHuEoFdmF9d
XaNYMFYwDQYDVR0OBAYEBAECAwQwDwYDVR0jBAgwBoAEAQIDBDASBgNVHRMBAf8E
CDAGAQH/AgEKMA8GA1UdJAEB/wQFMAOAAQAwDwYDVR0PAQH/BAUDAwf/gDAKBggq
hkjOPQQDAgNHADBEAiBMqieP2YMydkAXoagAHbzRRbJTxkd3SPHDiWhd9H9cUgIg
OWhAXP3wKuI/NEWzGS3jTVjNdkIZCc9cHOXxceA5Yrk=
-----END CERTIFICATE-----`

// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 1111638594 (0x42424242)
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	Validity
//	    Not Before: Feb 13 11:33:08 2018 GMT
//	    Not After : Dec 23 11:33:08 2027 GMT
//	Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (256 bit)
//	        pub:
//	            04:f1:bf:2d:e8:8c:66:40:e3:a8:d1:54:e0:42:49:
//	            02:cb:dd:47:08:85:c2:67:41:4c:eb:f7:87:cd:8d:
//	            a3:09:c8:18:cc:2e:30:53:16:32:aa:d5:9c:08:73:
//	            c6:76:fa:fa:3a:38:e9:34:35:9c:51:d1:ee:12:81:
//	            5d:98:5f:5d:5d
//	        ASN1 OID: prime256v1
//	        NIST CURVE: P-256
//	X509v3 extensions:
//	    X509v3 Subject Key Identifier:
//	        01:02:03:04
//	    X509v3 Authority Key Identifier:
//	        keyid:01:02:03:04
//	    X509v3 Basic Constraints: critical
//	        CA:TRUE, pathlen:10
//	    X509v3 Key Usage: critical
//	        Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//	    X509v3 Name Constraints:
//	        Permitted:
//	          DNS:.csr.pem
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	30:46:02:21:00:fd:11:41:d8:1f:2b:b5:49:8e:27:6e:70:93:
//	2c:f1:c2:e7:b0:a2:40:e2:c6:89:45:fc:99:a5:9b:dc:21:fb:
//	f6:02:21:00:b7:4f:98:bf:1f:dc:92:e7:db:7c:aa:33:7a:40:
//	36:1d:58:19:aa:96:3d:5e:5b:46:5f:47:f6:e3:7d:75:19:4f
const FakeIntermediateWithNameConstraintsCertPEM = `
-----BEGIN CERTIFICATE-----
MIICNjCCAdugAwIBAgIEQkJCQjAKBggqhkjOPQQDAjBxMQswCQYDVQQGEwJHQjEP
MA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdvb2ds
ZTEMMAoGA1UECxMDRW5nMSEwHwYDVQQDExhGYWtlQ2VydGlmaWNhdGVBdXRob3Jp
dHkwHhcNMTgwMjEzMTEzMzA4WhcNMjcxMjIzMTEzMzA4WjByMQswCQYDVQQGEwJH
QjEPMA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdv
b2dsZTEMMAoGA1UECxMDRW5nMSIwIAYDVQQDExlGYWtlSW50ZXJtZWRpYXRlQXV0
aG9yaXR5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8b8t6IxmQOOo0VTgQkkC
y91HCIXCZ0FM6/eHzY2jCcgYzC4wUxYyqtWcCHPGdvr6OjjpNDWcUdHuEoFdmF9d
XaNgMF4wDQYDVR0OBAYEBAECAwQwDwYDVR0jBAgwBoAEAQIDBDASBgNVHRMBAf8E
CDAGAQH/AgEKMA8GA1UdDwEB/wQFAwMH/4AwFwYDVR0eBBAwDqAMMAqCCC5jc3Iu
cGVtMAoGCCqGSM49BAMCA0kAMEYCIQD9EUHYHyu1SY4nbnCTLPHC57CiQOLGiUX8
maWb3CH79gIhALdPmL8f3JLn23yqM3pANh1YGaqWPV5bRl9H9uN9dRlP
-----END CERTIFICATE-----`

// Data:
//
//	Version: 3 (0x2)
//	Serial Number: 1111638594 (0x42424242)
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeCertificateAuthority
//	Validity
//	    Not Before: Feb 13 11:42:37 2018 GMT
//	    Not After : Dec 23 11:42:37 2027 GMT
//	Subject: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//	Subject Public Key Info:
//	    Public Key Algorithm: id-ecPublicKey
//	        Public-Key: (256 bit)
//	        pub:
//	            04:f1:bf:2d:e8:8c:66:40:e3:a8:d1:54:e0:42:49:
//	            02:cb:dd:47:08:85:c2:67:41:4c:eb:f7:87:cd:8d:
//	            a3:09:c8:18:cc:2e:30:53:16:32:aa:d5:9c:08:73:
//	            c6:76:fa:fa:3a:38:e9:34:35:9c:51:d1:ee:12:81:
//	            5d:98:5f:5d:5d
//	        ASN1 OID: prime256v1
//	        NIST CURVE: P-256
//	X509v3 extensions:
//	    X509v3 Subject Key Identifier:
//	        01:02:03:04
//	    X509v3 Authority Key Identifier:
//	        keyid:01:02:03:04
//
//	    X509v3 Basic Constraints: critical
//	        CA:TRUE, pathlen:10
//	    X509v3 Key Usage: critical
//	        Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Certificate Sign, CRL Sign, Encipher Only, Decipher Only
//	    X509v3 Name Constraints:
//	        Permitted:
//	          DNS:.xyzzy.pem
//
// Signature Algorithm: ecdsa-with-SHA256
//
//	30:45:02:20:3f:0a:40:60:b6:9e:ea:a5:cd:eb:e4:0e:7c:bc:
//	40:22:b2:e2:14:07:e8:ab:fa:4a:85:2a:41:18:20:f0:31:1a:
//	02:21:00:a4:64:91:6d:79:47:79:0f:16:06:62:a9:88:8b:92:
//	6d:40:fa:54:cb:c9:4f:bc:3f:53:27:e5:cd:12:16:53:7a
const FakeIntermediateWithInvalidNameConstraintsCertPEM = `
-----BEGIN CERTIFICATE-----
MIICNzCCAd2gAwIBAgIEQkJCQjAKBggqhkjOPQQDAjBxMQswCQYDVQQGEwJHQjEP
MA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdvb2ds
ZTEMMAoGA1UECxMDRW5nMSEwHwYDVQQDExhGYWtlQ2VydGlmaWNhdGVBdXRob3Jp
dHkwHhcNMTgwMjEzMTE0MjM3WhcNMjcxMjIzMTE0MjM3WjByMQswCQYDVQQGEwJH
QjEPMA0GA1UECBMGTG9uZG9uMQ8wDQYDVQQHEwZMb25kb24xDzANBgNVBAoTBkdv
b2dsZTEMMAoGA1UECxMDRW5nMSIwIAYDVQQDExlGYWtlSW50ZXJtZWRpYXRlQXV0
aG9yaXR5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8b8t6IxmQOOo0VTgQkkC
y91HCIXCZ0FM6/eHzY2jCcgYzC4wUxYyqtWcCHPGdvr6OjjpNDWcUdHuEoFdmF9d
XaNiMGAwDQYDVR0OBAYEBAECAwQwDwYDVR0jBAgwBoAEAQIDBDASBgNVHRMBAf8E
CDAGAQH/AgEKMA8GA1UdDwEB/wQFAwMH/4AwGQYDVR0eBBIwEKAOMAyCCi54eXp6
eS5wZW0wCgYIKoZIzj0EAwIDSAAwRQIgPwpAYLae6qXN6+QOfLxAIrLiFAfoq/pK
hSpBGCDwMRoCIQCkZJFteUd5DxYGYqmIi5JtQPpUy8lPvD9TJ+XNEhZTeg==
-----END CERTIFICATE-----`

// LeafCertPEM is a leaf cert signed by the key in:
//
//   - FakeIntermediateWithPolicyConstraintsCertPEM
//
//   - FakeIntermediateWithNameConstraintsCertPEM
//
//   - FakeIntermediateWithInvalidNameConstraintsCertPEM
//
//     Data:
//     Version: 3 (0x2)
//     Serial Number: 3735928559 (0xdeadbeef)
//     Signature Algorithm: ecdsa-with-SHA256
//     Issuer: C=GB, ST=London, L=London, O=Google, OU=Eng, CN=FakeIntermediateAuthority
//     Validity
//     Not Before: Feb 13 11:38:39 2018 GMT
//     Not After : Mar 28 11:38:39 2025 GMT
//     Subject: C=GB, ST=London, O=Google, OU=Eng, CN=leaf01.csr.pem
//     Subject Public Key Info:
//     Public Key Algorithm: id-ecPublicKey
//     Public-Key: (256 bit)
//     pub:
//     04:eb:37:4e:52:45:9c:46:d5:a8:b8:c5:ed:58:b9:
//     30:29:a6:70:8a:69:a0:26:5c:9e:2f:6e:b8:6b:23:
//     6c:84:e1:46:3a:98:36:82:44:a5:8a:17:8b:41:82:
//     32:f4:2d:e0:08:5b:7e:07:38:52:fc:47:56:28:27:
//     9b:ed:60:8b:ac
//     ASN1 OID: prime256v1
//     NIST CURVE: P-256
//     X509v3 extensions:
//     X509v3 Subject Key Identifier:
//     3F:B2:2F:41:FC:11:9A:D3:8D:A6:85:80:84:86:AE:7E:73:2E:69:5D
//     X509v3 Authority Key Identifier:
//     keyid:01:02:03:04
//     X509v3 Key Usage: critical
//     Digital Signature, Non Repudiation, Key Encipherment, Data Encipherment, Key Agreement, Encipher Only, Decipher Only
//     X509v3 Subject Alternative Name:
//     DNS:leaf01.csr.pem
//     Signature Algorithm: ecdsa-with-SHA256
//     30:46:02:21:00:b5:2a:f3:39:1e:06:b7:77:b2:ad:a8:83:1b:
//     83:38:64:5e:3a:25:51:e9:57:1f:00:53:72:db:08:11:65:3d:
//     f4:02:21:00:a1:4e:5d:b5:9a:8b:10:6e:15:a3:2a:bd:d9:80:
//     91:96:7c:1a:4f:8f:91:dc:44:9f:13:ff:57:f0:5e:ce:32:34
const LeafCertPEM = `
-----BEGIN CERTIFICATE-----
MIICGjCCAb+gAwIBAgIFAN6tvu8wCgYIKoZIzj0EAwIwcjELMAkGA1UEBhMCR0Ix
DzANBgNVBAgTBkxvbmRvbjEPMA0GA1UEBxMGTG9uZG9uMQ8wDQYDVQQKEwZHb29n
bGUxDDAKBgNVBAsTA0VuZzEiMCAGA1UEAxMZRmFrZUludGVybWVkaWF0ZUF1dGhv
cml0eTAeFw0xODAyMTMxMTM4MzlaFw0yNTAzMjgxMTM4MzlaMFYxCzAJBgNVBAYT
AkdCMQ8wDQYDVQQIDAZMb25kb24xDzANBgNVBAoMBkdvb2dsZTEMMAoGA1UECwwD
RW5nMRcwFQYDVQQDDA5sZWFmMDEuY3NyLnBlbTBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABOs3TlJFnEbVqLjF7Vi5MCmmcIppoCZcni9uuGsjbIThRjqYNoJEpYoX
i0GCMvQt4Ahbfgc4UvxHVignm+1gi6yjXjBcMB0GA1UdDgQWBBQ/si9B/BGa042m
hYCEhq5+cy5pXTAPBgNVHSMECDAGgAQBAgMEMA8GA1UdDwEB/wQFAwMH+YAwGQYD
VR0RBBIwEIIObGVhZjAxLmNzci5wZW0wCgYIKoZIzj0EAwIDSQAwRgIhALUq8zke
Brd3sq2ogxuDOGReOiVR6VcfAFNy2wgRZT30AiEAoU5dtZqLEG4Voyq92YCRlnwa
T4+R3ESfE/9X8F7OMjQ=
-----END CERTIFICATE-----`

// Issuance chain 6
// ================
// The next section holds a real world intermediate and leaf cert.

// RealPrecertIntermediatePEM is the intermediate issuer for
// RealPrecertWithEKUPEM, below.
//
// Certificate:
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number:
//	        01:e3:b4:9d:77:cd:f4:0c:06:19:16:b6:e3
//	    Signature Algorithm: sha256WithRSAEncryption
//	    Issuer: OU = GlobalSign Root CA - R2, O = GlobalSign, CN = GlobalSign
//	    Validity
//	        Not Before: Jun 15 00:00:42 2017 GMT
//	        Not After : Dec 15 00:00:42 2021 GMT
//	    Subject: C = US, O = Google Trust Services, CN = GTS CA 1D2
//	    Subject Public Key Info:
//	        Public Key Algorithm: rsaEncryption
//	            RSA Public-Key: (2048 bit)
//	            Modulus:
//	                00:b2:d9:7b:e1:e1:d7:3f:1c:91:72:ff:f9:10:cd:
//	                87:15:79:74:b7:3e:47:8b:b2:61:55:fd:0c:36:c6:
//	                7e:77:42:3a:b2:fa:52:5b:0b:71:81:d6:4d:d5:e9:
//	                2b:24:4d:23:5e:8b:2b:72:5f:21:55:b5:29:ef:44:
//	                cb:eb:82:52:ab:3e:27:a4:92:49:41:4a:de:a8:dd:
//	                31:e0:3c:df:6d:7a:4d:2d:d6:6d:09:b0:0e:e3:61:
//	                f2:b2:fe:90:6c:5a:7b:10:64:49:b4:0b:3c:08:f2:
//	                ea:79:0c:6c:a6:1a:89:6a:56:32:a0:29:a2:30:82:
//	                8f:81:51:0c:f3:a2:b9:d9:75:b9:22:9e:27:14:ba:
//	                4a:2f:2c:63:58:87:f1:5d:10:e6:5f:91:bb:b9:5b:
//	                cc:47:e2:1e:75:b6:8c:8f:cc:75:5d:57:05:e7:82:
//	                c6:84:0e:74:72:2a:cb:3b:55:f5:6e:70:eb:66:69:
//	                c3:24:bb:38:93:35:9b:68:61:2f:9b:d6:ae:a6:77:
//	                72:7c:71:48:58:33:10:af:e9:80:82:1d:b5:07:40:
//	                1b:f6:3d:ec:a2:ad:47:9d:b4:94:29:34:b3:8c:2f:
//	                cd:25:03:58:35:c0:25:a4:55:5f:e1:b3:07:56:3d:
//	                c8:d0:63:b8:20:fb:8c:1d:43:2c:f8:f9:a9:d5:ec:
//	                6f:97
//	            Exponent: 65537 (0x10001)
//	    X509v3 extensions:
//	        X509v3 Key Usage: critical
//	            Digital Signature, Certificate Sign, CRL Sign
//	        X509v3 Extended Key Usage:
//	            TLS Web Server Authentication, TLS Web Client Authentication
//	        X509v3 Basic Constraints: critical
//	            CA:TRUE, pathlen:0
//	        X509v3 Subject Key Identifier:
//	            B1:DD:32:5D:E8:B7:37:72:D2:CE:5C:CE:26:FE:47:79:E2:01:08:E9
//	        X509v3 Authority Key Identifier:
//	            keyid:9B:E2:07:57:67:1C:1E:C0:6A:06:DE:59:B4:9A:2D:DF:DC:19:86:2E
//
//	        Authority Information Access:
//	            OCSP - URI:http://ocsp.pki.goog/gsr2
//
//	        X509v3 CRL Distribution Points:
//
//	            Full Name:
//	              URI:http://crl.pki.goog/gsr2/gsr2.crl
//
//	        X509v3 Certificate Policies:
//	            Policy: 2.23.140.1.2.1
//	              CPS: https://pki.goog/repository/
//
//	Signature Algorithm: sha256WithRSAEncryption
//	     71:4a:c4:c3:23:ae:f7:e3:b2:02:79:8c:13:e8:53:8e:80:c5:
//	     f0:e3:ef:71:60:a9:a9:7b:34:65:85:34:bd:47:3b:03:57:16:
//	     00:99:48:3a:e0:e0:f0:ea:cd:b6:48:3c:d5:ab:72:f0:d0:1b:
//	     cb:64:2d:3b:0d:74:68:d7:74:88:31:7c:6a:ba:0e:f0:8c:4d:
//	     78:ce:da:10:f4:8a:96:45:97:a9:97:ad:c5:35:1a:18:64:e8:
//	     93:b6:0d:9d:1f:b9:5e:1d:80:ea:e7:5b:9c:8e:ae:0e:a6:84:
//	     d2:d1:17:ce:b3:fb:f6:81:4f:3c:e6:68:9f:cf:f1:a6:76:c5:
//	     7d:a7:f3:dd:7d:58:0f:e0:f6:61:01:1c:51:8e:76:33:2b:48:
//	     9d:5c:81:51:72:08:17:ba:fd:01:d3:ee:46:f9:f4:b2:68:40:
//	     99:31:01:6c:4f:1b:c6:56:eb:81:73:d2:79:52:05:92:26:5b:
//	     71:cd:9d:c4:d2:ce:23:77:0f:41:7a:69:5e:21:25:c6:f8:b7:
//	     ff:7a:f7:47:de:c2:00:7b:9c:5a:45:9c:2a:4e:46:90:d9:75:
//	     2c:d8:ff:8c:ee:cc:dc:69:eb:6c:e6:15:d0:a3:ff:48:0b:ac:
//	     55:df:df:25:9d:42:b6:51:a3:66:95:60:c5:d0:22:e7:22:7a:
//	     51:a5:cc:87
const RealPrecertIntermediatePEM = `
-----BEGIN CERTIFICATE-----
MIIESjCCAzKgAwIBAgINAeO0nXfN9AwGGRa24zANBgkqhkiG9w0BAQsFADBMMSAw
HgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMjETMBEGA1UEChMKR2xvYmFs
U2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjAeFw0xNzA2MTUwMDAwNDJaFw0yMTEy
MTUwMDAwNDJaMEIxCzAJBgNVBAYTAlVTMR4wHAYDVQQKExVHb29nbGUgVHJ1c3Qg
U2VydmljZXMxEzARBgNVBAMTCkdUUyBDQSAxRDIwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQCy2Xvh4dc/HJFy//kQzYcVeXS3PkeLsmFV/Qw2xn53Qjqy
+lJbC3GB1k3V6SskTSNeiytyXyFVtSnvRMvrglKrPiekkklBSt6o3THgPN9tek0t
1m0JsA7jYfKy/pBsWnsQZEm0CzwI8up5DGymGolqVjKgKaIwgo+BUQzzornZdbki
nicUukovLGNYh/FdEOZfkbu5W8xH4h51toyPzHVdVwXngsaEDnRyKss7VfVucOtm
acMkuziTNZtoYS+b1q6md3J8cUhYMxCv6YCCHbUHQBv2PeyirUedtJQpNLOML80l
A1g1wCWkVV/hswdWPcjQY7gg+4wdQyz4+anV7G+XAgMBAAGjggEzMIIBLzAOBgNV
HQ8BAf8EBAMCAYYwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMBIGA1Ud
EwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFLHdMl3otzdy0s5czib+R3niAQjpMB8G
A1UdIwQYMBaAFJviB1dnHB7AagbeWbSaLd/cGYYuMDUGCCsGAQUFBwEBBCkwJzAl
BggrBgEFBQcwAYYZaHR0cDovL29jc3AucGtpLmdvb2cvZ3NyMjAyBgNVHR8EKzAp
MCegJaAjhiFodHRwOi8vY3JsLnBraS5nb29nL2dzcjIvZ3NyMi5jcmwwPwYDVR0g
BDgwNjA0BgZngQwBAgEwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly9wa2kuZ29vZy9y
ZXBvc2l0b3J5LzANBgkqhkiG9w0BAQsFAAOCAQEAcUrEwyOu9+OyAnmME+hTjoDF
8OPvcWCpqXs0ZYU0vUc7A1cWAJlIOuDg8OrNtkg81aty8NAby2QtOw10aNd0iDF8
aroO8IxNeM7aEPSKlkWXqZetxTUaGGTok7YNnR+5Xh2A6udbnI6uDqaE0tEXzrP7
9oFPPOZon8/xpnbFfafz3X1YD+D2YQEcUY52MytInVyBUXIIF7r9AdPuRvn0smhA
mTEBbE8bxlbrgXPSeVIFkiZbcc2dxNLOI3cPQXppXiElxvi3/3r3R97CAHucWkWc
Kk5GkNl1LNj/jO7M3GnrbOYV0KP/SAusVd/fJZ1CtlGjZpVgxdAi5yJ6UaXMhw==
-----END CERTIFICATE-----
`

// RealPrecertWithEKUPEM is an actual precertificate containing a valid EKU.
//
// Certificate:
//
//	Data:
//	    Version: 3 (0x2)
//	    Serial Number:
//	        9c:a4:07:e2:25:f9:7c:c2:0a:00:00:00:00:20:6e:e5
//	    Signature Algorithm: sha256WithRSAEncryption
//	    Issuer: C = US, O = Google Trust Services, CN = GTS CA 1D2
//	    Validity
//	        Not Before: Mar 23 12:23:44 2020 GMT
//	        Not After : Jun 21 12:23:44 2020 GMT
//	    Subject: CN = certificate.transparency.dev
//	    Subject Public Key Info:
//	        Public Key Algorithm: rsaEncryption
//	            RSA Public-Key: (2048 bit)
//	            Modulus:
//	                00:a8:7e:59:c0:e5:3b:da:3c:bf:04:51:91:eb:9f:
//	                6c:1b:cf:9f:90:dc:22:89:1c:b5:98:24:69:2e:26:
//	                2d:61:92:04:0f:2e:f1:da:ec:ea:3a:d9:cc:3a:82:
//	                e2:b8:3a:7d:6c:79:79:f7:36:c5:52:a4:bb:46:1d:
//	                2f:0b:6c:5f:00:31:af:24:e9:4a:1b:32:63:1a:b5:
//	                c3:28:9c:a7:0a:b5:73:e2:c1:a7:b5:1e:11:ae:cd:
//	                19:79:0c:62:06:cf:80:f0:ed:e2:72:82:bb:b4:84:
//	                0e:9d:c9:7d:3b:fb:4e:05:49:3a:14:0f:86:92:01:
//	                49:52:2c:cc:a0:e1:ef:86:fe:18:00:83:69:6c:90:
//	                c6:7b:a9:42:df:57:9c:7b:61:06:80:23:b2:5f:95:
//	                95:1e:9b:34:6f:ab:a3:21:1b:2b:8e:9f:34:4f:ec:
//	                e8:9a:48:74:81:2f:9b:12:67:54:a1:46:76:96:9a:
//	                1e:9d:c3:ee:bf:6a:e8:49:72:57:28:b1:12:c4:ca:
//	                41:84:96:f7:32:4a:4a:9e:59:2d:48:3e:ac:29:0c:
//	                f4:f4:03:28:33:1a:73:10:48:29:68:12:e3:f9:7e:
//	                f4:5f:01:54:b0:73:c6:a8:72:b6:84:54:05:23:36:
//	                b6:db:3f:d8:e5:27:89:4c:dc:bb:b1:c9:9e:e7:7e:
//	                b0:b5
//	            Exponent: 65537 (0x10001)
//	    X509v3 extensions:
//	        X509v3 Key Usage: critical
//	            Digital Signature, Key Encipherment
//	        X509v3 Extended Key Usage:
//	            TLS Web Server Authentication
//	        X509v3 Basic Constraints: critical
//	            CA:FALSE
//	        X509v3 Subject Key Identifier:
//	            B8:E0:AF:4F:7C:48:F3:FF:EB:FC:5E:A5:34:36:2D:56:54:AC:97:6B
//	        X509v3 Authority Key Identifier:
//	            keyid:B1:DD:32:5D:E8:B7:37:72:D2:CE:5C:CE:26:FE:47:79:E2:01:08:E9
//
//	        Authority Information Access:
//	            OCSP - URI:http://ocsp.pki.goog/gts1d2
//	            CA Issuers - URI:http://pki.goog/gsr2/GTS1D2.crt
//
//	        X509v3 Subject Alternative Name:
//	            DNS:certificate.transparency.dev
//	        X509v3 Certificate Policies:
//	            Policy: 2.23.140.1.2.1
//	            Policy: 1.3.6.1.4.1.11129.2.5.3
//
//	        X509v3 CRL Distribution Points:
//
//	            Full Name:
//	              URI:http://crl.pki.goog/GTS1D2.crl
//
//	        CT Precertificate Poison: critical
//	            NULL
//	Signature Algorithm: sha256WithRSAEncryption
//	     51:fe:93:53:7a:e1:6d:34:ce:a2:1d:4d:32:c5:39:a5:e8:1e:
//	     ee:97:56:33:84:5a:5e:5c:be:13:64:92:66:df:a7:79:82:c8:
//	     35:c6:4d:8f:ff:da:a1:cc:4d:70:b0:a7:1c:73:69:d5:08:ea:
//	     53:f4:8e:73:27:5a:9d:5a:c7:39:0a:19:dd:51:21:94:3c:31:
//	     b5:cd:06:2d:50:bf:90:09:3e:62:ca:a3:bf:f2:74:9d:2b:33:
//	     38:e9:9f:f1:b7:2f:e2:3c:e4:8a:d4:63:57:c7:bd:27:fd:94:
//	     15:c5:03:82:95:35:79:d6:84:0f:90:01:47:53:af:ed:12:d6:
//	     9c:63:04:1b:06:83:87:83:a1:34:f0:05:d8:8b:c6:b9:39:ce:
//	     9c:32:ac:bf:04:d5:8d:b8:2f:ee:61:55:b9:f3:b9:b8:93:c7:
//	     6d:9c:39:68:b4:39:d8:67:5d:cb:5b:bd:d5:a1:b8:d9:18:16:
//	     7c:f3:ff:7a:77:d9:cc:68:f3:c8:ee:b4:52:06:37:6c:8e:23:
//	     69:1c:49:81:1c:08:26:80:a1:05:8b:ed:f5:dc:33:c6:84:7a:
//	     e3:ef:2f:c3:22:02:a0:33:8d:48:61:8a:98:27:34:e8:75:5d:
//	     eb:56:93:a3:be:2e:c5:04:ab:d6:88:cc:53:c6:9c:db:9f:aa:
//	     5d:eb:c6:82
const RealPrecertWithEKUPEM = `
-----BEGIN CERTIFICATE-----
MIIEZTCCA02gAwIBAgIRAJykB+Il+XzCCgAAAAAgbuUwDQYJKoZIhvcNAQELBQAw
QjELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFUdvb2dsZSBUcnVzdCBTZXJ2aWNlczET
MBEGA1UEAxMKR1RTIENBIDFEMjAeFw0yMDAzMjMxMjIzNDRaFw0yMDA2MjExMjIz
NDRaMCcxJTAjBgNVBAMTHGNlcnRpZmljYXRlLnRyYW5zcGFyZW5jeS5kZXYwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCoflnA5TvaPL8EUZHrn2wbz5+Q
3CKJHLWYJGkuJi1hkgQPLvHa7Oo62cw6guK4On1seXn3NsVSpLtGHS8LbF8AMa8k
6UobMmMatcMonKcKtXPiwae1HhGuzRl5DGIGz4Dw7eJygru0hA6dyX07+04FSToU
D4aSAUlSLMyg4e+G/hgAg2lskMZ7qULfV5x7YQaAI7JflZUemzRvq6MhGyuOnzRP
7OiaSHSBL5sSZ1ShRnaWmh6dw+6/auhJclcosRLEykGElvcySkqeWS1IPqwpDPT0
AygzGnMQSCloEuP5fvRfAVSwc8aocraEVAUjNrbbP9jlJ4lM3LuxyZ7nfrC1AgMB
AAGjggFvMIIBazAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwEw
DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUuOCvT3xI8//r/F6lNDYtVlSsl2swHwYD
VR0jBBgwFoAUsd0yXei3N3LSzlzOJv5HeeIBCOkwZAYIKwYBBQUHAQEEWDBWMCcG
CCsGAQUFBzABhhtodHRwOi8vb2NzcC5wa2kuZ29vZy9ndHMxZDIwKwYIKwYBBQUH
MAKGH2h0dHA6Ly9wa2kuZ29vZy9nc3IyL0dUUzFEMi5jcnQwJwYDVR0RBCAwHoIc
Y2VydGlmaWNhdGUudHJhbnNwYXJlbmN5LmRldjAhBgNVHSAEGjAYMAgGBmeBDAEC
ATAMBgorBgEEAdZ5AgUDMC8GA1UdHwQoMCYwJKAioCCGHmh0dHA6Ly9jcmwucGtp
Lmdvb2cvR1RTMUQyLmNybDATBgorBgEEAdZ5AgQDAQH/BAIFADANBgkqhkiG9w0B
AQsFAAOCAQEAUf6TU3rhbTTOoh1NMsU5pege7pdWM4RaXly+E2SSZt+neYLINcZN
j//aocxNcLCnHHNp1QjqU/SOcydanVrHOQoZ3VEhlDwxtc0GLVC/kAk+Ysqjv/J0
nSszOOmf8bcv4jzkitRjV8e9J/2UFcUDgpU1edaED5ABR1Ov7RLWnGMEGwaDh4Oh
NPAF2IvGuTnOnDKsvwTVjbgv7mFVufO5uJPHbZw5aLQ52Gddy1u91aG42RgWfPP/
enfZzGjzyO60UgY3bI4jaRxJgRwIJoChBYvt9dwzxoR64+8vwyICoDONSGGKmCc0
6HVd61aTo74uxQSr1ojMU8ac25+qXevGgg==
-----END CERTIFICATE-----
`
