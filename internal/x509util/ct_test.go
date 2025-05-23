// Copyright 2024 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package x509util

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"
)

var pemPrivateKey = testingKey(`
-----BEGIN RSA TESTING KEY-----
MIICXAIBAAKBgQCxoeCUW5KJxNPxMp+KmCxKLc1Zv9Ny+4CFqcUXVUYH69L3mQ7v
IWrJ9GBfcaA7BPQqUlWxWM+OCEQZH1EZNIuqRMNQVuIGCbz5UQ8w6tS0gcgdeGX7
J7jgCQ4RK3F/PuCM38QBLaHx988qG8NMc6VKErBjctCXFHQt14lerd5KpQIDAQAB
AoGAYrf6Hbk+mT5AI33k2Jt1kcweodBP7UkExkPxeuQzRVe0KVJw0EkcFhywKpr1
V5eLMrILWcJnpyHE5slWwtFHBG6a5fLaNtsBBtcAIfqTQ0Vfj5c6SzVaJv0Z5rOd
7gQF6isy3t3w9IF3We9wXQKzT6q5ypPGdm6fciKQ8RnzREkCQQDZwppKATqQ41/R
vhSj90fFifrGE6aVKC1hgSpxGQa4oIdsYYHwMzyhBmWW9Xv/R+fPyr8ZwPxp2c12
33QwOLPLAkEA0NNUb+z4ebVVHyvSwF5jhfJxigim+s49KuzJ1+A2RaSApGyBZiwS
rWvWkB471POAKUYt5ykIWVZ83zcceQiNTwJBAMJUFQZX5GDqWFc/zwGoKkeR49Yi
MTXIvf7Wmv6E++eFcnT461FlGAUHRV+bQQXGsItR/opIG7mGogIkVXa3E1MCQARX
AAA7eoZ9AEHflUeuLn9QJI/r0hyQQLEtrpwv6rDT1GCWaLII5HJ6NUFVf4TTcqxo
6vdM4QGKTJoO+SaCyP0CQFdpcxSAuzpFcKv0IlJ8XzS/cy+mweCMwyJ1PFEc4FX6
wg/HcAJWY60xZTJDFN+Qfx8ZQvBEin6c2/h+zZi5IVY=
-----END RSA TESTING KEY-----
`)

var testPrivateKey *rsa.PrivateKey

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func init() {
	block, _ := pem.Decode([]byte(pemPrivateKey))

	var err error
	if testPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		panic("Failed to parse private key: " + err.Error())
	}
}

func makeCert(t *testing.T, template, issuer *x509.Certificate) *x509.Certificate {
	t.Helper()
	certData, err := x509.CreateCertificate(rand.Reader, template, issuer, &testPrivateKey.PublicKey, testPrivateKey)
	if err != nil {
		t.Fatalf("failed to create pre-cert: %v", err)
	}
	cert, err := x509.ParseCertificate(certData)
	if err != nil {
		t.Fatalf("failed to re-parse pre-cert: %v", err)
	}
	return cert
}

func TestBuildPrecertTBS(t *testing.T) {
	poisonExt := pkix.Extension{Id: oidExtensionCTPoison, Critical: true, Value: asn1.NullBytes}
	// TODO(phboneff): check Critical and value are ok.
	ctExt := pkix.Extension{Id: oidExtensionKeyUsageCertificateTransparency}
	preIssuerKeyID := []byte{0x19, 0x09, 0x19, 0x70}
	issuerKeyID := []byte{0x07, 0x07, 0x20, 0x07}
	preCertTemplate := x509.Certificate{
		Version:         3,
		SerialNumber:    big.NewInt(123),
		Issuer:          pkix.Name{CommonName: "precert Issuer"},
		Subject:         pkix.Name{CommonName: "precert subject"},
		NotBefore:       time.Now(),
		NotAfter:        time.Now().Add(3 * time.Hour),
		ExtraExtensions: []pkix.Extension{poisonExt},
		AuthorityKeyId:  preIssuerKeyID,
	}
	preIssuerTemplate := x509.Certificate{
		Version:         3,
		SerialNumber:    big.NewInt(1234),
		Issuer:          pkix.Name{CommonName: "real Issuer"},
		Subject:         pkix.Name{CommonName: "precert Issuer"},
		NotBefore:       time.Now(),
		NotAfter:        time.Now().Add(3 * time.Hour),
		ExtraExtensions: []pkix.Extension{ctExt},
		AuthorityKeyId:  issuerKeyID,
		SubjectKeyId:    preIssuerKeyID,
	}
	actualIssuerTemplate := x509.Certificate{
		Version:      3,
		SerialNumber: big.NewInt(12345),
		Issuer:       pkix.Name{CommonName: "real Issuer"},
		Subject:      pkix.Name{CommonName: "real Issuer"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(3 * time.Hour),
		SubjectKeyId: issuerKeyID,
	}
	preCertWithAKI := makeCert(t, &preCertTemplate, &preIssuerTemplate)
	preIssuerWithAKI := makeCert(t, &preIssuerTemplate, &actualIssuerTemplate)

	preIssuerTemplate.AuthorityKeyId = nil
	actualIssuerTemplate.SubjectKeyId = nil
	preIssuerWithoutAKI := makeCert(t, &preIssuerTemplate, &actualIssuerTemplate)

	preCertTemplate.AuthorityKeyId = nil
	preIssuerTemplate.SubjectKeyId = nil
	preCertWithoutAKI := makeCert(t, &preCertTemplate, &preIssuerTemplate)

	preIssuerTemplate.ExtraExtensions = nil
	invalidPreIssuer := makeCert(t, &preIssuerTemplate, &actualIssuerTemplate)

	akiPrefix := []byte{0x30, 0x06, 0x80, 0x04} // SEQUENCE { [0] { ... } }
	var tests = []struct {
		name      string
		tbs       *x509.Certificate
		preIssuer *x509.Certificate
		wantAKI   []byte
		wantErr   bool
	}{
		{
			name:    "no-preIssuer-provided",
			tbs:     preCertWithAKI,
			wantAKI: append(akiPrefix, preIssuerKeyID...),
		},
		{
			name:      "both-with-AKI",
			tbs:       preCertWithAKI,
			preIssuer: preIssuerWithAKI,
			wantAKI:   append(akiPrefix, issuerKeyID...),
		},
		{
			name:      "invalid-preIssuer",
			tbs:       preCertWithAKI,
			preIssuer: invalidPreIssuer,
			wantErr:   true,
		},
		{
			name:      "both-without-AKI",
			tbs:       preCertWithoutAKI,
			preIssuer: preIssuerWithoutAKI,
		},
		{
			name:      "precert-with-preIssuer-without-AKI",
			tbs:       preCertWithAKI,
			preIssuer: preIssuerWithoutAKI,
		},
		{
			name:      "precert-without-preIssuer-with-AKI",
			tbs:       preCertWithoutAKI,
			preIssuer: preIssuerWithAKI,
			wantAKI:   append(akiPrefix, issuerKeyID...),
		},
	}
	for _, test := range tests {
		got, err := BuildPrecertTBS(test.tbs.RawTBSCertificate, test.preIssuer)
		if err != nil {
			if !test.wantErr {
				t.Errorf("BuildPrecertTBS(%s)=nil,%q; want _,nil", test.name, err)
			}
			continue
		}
		if test.wantErr {
			t.Errorf("BuildPrecertTBS(%s)=_,nil; want _,non-nil", test.name)
		}

		var tbs tbsCertificate
		if rest, err := asn1.Unmarshal(got, &tbs); err != nil {
			t.Errorf("BuildPrecertTBS(%s) gave unparsable TBS: %v", test.name, err)
			continue
		} else if len(rest) > 0 {
			t.Errorf("BuildPrecertTBS(%s) gave extra data in DER", test.name)
		}
		if test.preIssuer != nil {
			if got, want := tbs.Issuer.FullBytes, test.preIssuer.RawIssuer; !bytes.Equal(got, want) {
				t.Errorf("BuildPrecertTBS(%s).Issuer=%x, want %x", test.name, got, want)
			}
		}
		var gotAKI []byte
		for _, ext := range tbs.Extensions {
			if ext.Id.Equal(oidExtensionAuthorityKeyId) {
				gotAKI = ext.Value
				break
			}
		}
		if gotAKI != nil {
			if test.wantAKI != nil {
				if !reflect.DeepEqual(gotAKI, test.wantAKI) {
					t.Errorf("BuildPrecertTBS(%s).Extensions[AKI]=%+v, want %+v", test.name, gotAKI, test.wantAKI)
				}
			} else {
				t.Errorf("BuildPrecertTBS(%s).Extensions[AKI]=%+v, want nil", test.name, gotAKI)
			}
		} else if test.wantAKI != nil {
			t.Errorf("BuildPrecertTBS(%s).Extensions[AKI]=nil, want %+v", test.name, test.wantAKI)
		}
	}
}

const (
	tbsNoPoison = "30820245a003020102020842822a5b866fbfeb300d06092a864886f70d01010b" +
		"05003071310b3009060355040613024742310f300d060355040813064c6f6e64" +
		"6f6e310f300d060355040713064c6f6e646f6e310f300d060355040a1306476f" +
		"6f676c65310c300a060355040b1303456e673121301f0603550403131846616b" +
		"654365727469666963617465417574686f72697479301e170d31363037313731" +
		"31313534305a170d3139303331393131313534305a3066310b30090603550406" +
		"130255533113301106035504080c0a43616c69666f726e696131163014060355" +
		"04070c0d4d6f756e7461696e205669657731133011060355040a0c0a476f6f67" +
		"6c6520496e633115301306035504030c0c2a2e676f6f676c652e636f6d305930" +
		"1306072a8648ce3d020106082a8648ce3d03010703420004c4093984f5158d12" +
		"54b2029cf901e26d3547d40dd011616609351dcb121495b23fff35bd228e4dfc" +
		"38502d22d6981ecaa023afa4967e32d1825f3157fb28ff37a381ce3081cb301d" +
		"0603551d250416301406082b0601050507030106082b06010505070302306806" +
		"082b06010505070101045c305a302b06082b06010505073002861f687474703a" +
		"2f2f706b692e676f6f676c652e636f6d2f47494147322e637274302b06082b06" +
		"010505073001861f687474703a2f2f636c69656e7473312e676f6f676c652e63" +
		"6f6d2f6f637370301d0603551d0e04160414dbf46e63eee2dcbebf38604f9831" +
		"d06444f163d830210603551d20041a3018300c060a2b06010401d67902050130" +
		"08060667810c010202"
	tbsPoisonFirst = "3082025aa003020102020842822a5b866fbfeb300d06092a864886f70d01010b" +
		"05003071310b3009060355040613024742310f300d060355040813064c6f6e64" +
		"6f6e310f300d060355040713064c6f6e646f6e310f300d060355040a1306476f" +
		"6f676c65310c300a060355040b1303456e673121301f0603550403131846616b" +
		"654365727469666963617465417574686f72697479301e170d31363037313731" +
		"31313534305a170d3139303331393131313534305a3066310b30090603550406" +
		"130255533113301106035504080c0a43616c69666f726e696131163014060355" +
		"04070c0d4d6f756e7461696e205669657731133011060355040a0c0a476f6f67" +
		"6c6520496e633115301306035504030c0c2a2e676f6f676c652e636f6d305930" +
		"1306072a8648ce3d020106082a8648ce3d03010703420004c4093984f5158d12" +
		"54b2029cf901e26d3547d40dd011616609351dcb121495b23fff35bd228e4dfc" +
		"38502d22d6981ecaa023afa4967e32d1825f3157fb28ff37a381e33081e03013" +
		"060a2b06010401d6790204030101ff04020500301d0603551d25041630140608" +
		"2b0601050507030106082b06010505070302306806082b06010505070101045c" +
		"305a302b06082b06010505073002861f687474703a2f2f706b692e676f6f676c" +
		"652e636f6d2f47494147322e637274302b06082b06010505073001861f687474" +
		"703a2f2f636c69656e7473312e676f6f676c652e636f6d2f6f637370301d0603" +
		"551d0e04160414dbf46e63eee2dcbebf38604f9831d06444f163d83021060355" +
		"1d20041a3018300c060a2b06010401d6790205013008060667810c010202"
	tbsPoisonLast = "3082025aa003020102020842822a5b866fbfeb300d06092a864886f70d01010b" +
		"05003071310b3009060355040613024742310f300d060355040813064c6f6e64" +
		"6f6e310f300d060355040713064c6f6e646f6e310f300d060355040a1306476f" +
		"6f676c65310c300a060355040b1303456e673121301f0603550403131846616b" +
		"654365727469666963617465417574686f72697479301e170d31363037313731" +
		"31313534305a170d3139303331393131313534305a3066310b30090603550406" +
		"130255533113301106035504080c0a43616c69666f726e696131163014060355" +
		"04070c0d4d6f756e7461696e205669657731133011060355040a0c0a476f6f67" +
		"6c6520496e633115301306035504030c0c2a2e676f6f676c652e636f6d305930" +
		"1306072a8648ce3d020106082a8648ce3d03010703420004c4093984f5158d12" +
		"54b2029cf901e26d3547d40dd011616609351dcb121495b23fff35bd228e4dfc" +
		"38502d22d6981ecaa023afa4967e32d1825f3157fb28ff37a381e33081e0301d" +
		"0603551d250416301406082b0601050507030106082b06010505070302306806" +
		"082b06010505070101045c305a302b06082b06010505073002861f687474703a" +
		"2f2f706b692e676f6f676c652e636f6d2f47494147322e637274302b06082b06" +
		"010505073001861f687474703a2f2f636c69656e7473312e676f6f676c652e63" +
		"6f6d2f6f637370301d0603551d0e04160414dbf46e63eee2dcbebf38604f9831" +
		"d06444f163d830210603551d20041a3018300c060a2b06010401d67902050130" +
		"08060667810c0102023013060a2b06010401d6790204030101ff04020500"
	tbsPoisonMiddle = "3082025aa003020102020842822a5b866fbfeb300d06092a864886f70d01010b" +
		"05003071310b3009060355040613024742310f300d060355040813064c6f6e64" +
		"6f6e310f300d060355040713064c6f6e646f6e310f300d060355040a1306476f" +
		"6f676c65310c300a060355040b1303456e673121301f0603550403131846616b" +
		"654365727469666963617465417574686f72697479301e170d31363037313731" +
		"31313534305a170d3139303331393131313534305a3066310b30090603550406" +
		"130255533113301106035504080c0a43616c69666f726e696131163014060355" +
		"04070c0d4d6f756e7461696e205669657731133011060355040a0c0a476f6f67" +
		"6c6520496e633115301306035504030c0c2a2e676f6f676c652e636f6d305930" +
		"1306072a8648ce3d020106082a8648ce3d03010703420004c4093984f5158d12" +
		"54b2029cf901e26d3547d40dd011616609351dcb121495b23fff35bd228e4dfc" +
		"38502d22d6981ecaa023afa4967e32d1825f3157fb28ff37a381e33081e0301d" +
		"0603551d250416301406082b0601050507030106082b06010505070302306806" +
		"082b06010505070101045c305a302b06082b06010505073002861f687474703a" +
		"2f2f706b692e676f6f676c652e636f6d2f47494147322e637274302b06082b06" +
		"010505073001861f687474703a2f2f636c69656e7473312e676f6f676c652e63" +
		"6f6d2f6f6373703013060a2b06010401d6790204030101ff04020500301d0603" +
		"551d0e04160414dbf46e63eee2dcbebf38604f9831d06444f163d83021060355" +
		"1d20041a3018300c060a2b06010401d6790205013008060667810c010202"
	tbsPoisonTwice = "3082026fa003020102020842822a5b866fbfeb300d06092a864886f70d01010b" +
		"05003071310b3009060355040613024742310f300d060355040813064c6f6e64" +
		"6f6e310f300d060355040713064c6f6e646f6e310f300d060355040a1306476f" +
		"6f676c65310c300a060355040b1303456e673121301f0603550403131846616b" +
		"654365727469666963617465417574686f72697479301e170d31363037313731" +
		"31313534305a170d3139303331393131313534305a3066310b30090603550406" +
		"130255533113301106035504080c0a43616c69666f726e696131163014060355" +
		"04070c0d4d6f756e7461696e205669657731133011060355040a0c0a476f6f67" +
		"6c6520496e633115301306035504030c0c2a2e676f6f676c652e636f6d305930" +
		"1306072a8648ce3d020106082a8648ce3d03010703420004c4093984f5158d12" +
		"54b2029cf901e26d3547d40dd011616609351dcb121495b23fff35bd228e4dfc" +
		"38502d22d6981ecaa023afa4967e32d1825f3157fb28ff37a381f83081f5301d" +
		"0603551d250416301406082b0601050507030106082b06010505070302306806" +
		"082b06010505070101045c305a302b06082b06010505073002861f687474703a" +
		"2f2f706b692e676f6f676c652e636f6d2f47494147322e637274302b06082b06" +
		"010505073001861f687474703a2f2f636c69656e7473312e676f6f676c652e63" +
		"6f6d2f6f6373703013060a2b06010401d6790204030101ff04020500301d0603" +
		"551d0e04160414dbf46e63eee2dcbebf38604f9831d06444f163d83013060a2b" +
		"06010401d6790204030101ff0402050030210603551d20041a3018300c060a2b" +
		"06010401d6790205013008060667810c010202"
)

func TestRemoveCTPoison(t *testing.T) {
	var tests = []struct {
		name   string // for human consumption
		tbs    string // hex encoded
		want   string // hex encoded
		errstr string
	}{
		{name: "invalid-der", tbs: "01020304", errstr: "failed to parse"},
		{name: "trailing-data", tbs: tbsPoisonMiddle + "01020304", errstr: "trailing data"},
		{name: "no-poison-ext", tbs: tbsNoPoison, errstr: "no extension of specified type present"},
		{name: "two-poison-exts", tbs: tbsPoisonTwice, errstr: "multiple extensions of specified type present"},
		{name: "poison-first", tbs: tbsPoisonFirst, want: tbsNoPoison},
		{name: "poison-last", tbs: tbsPoisonLast, want: tbsNoPoison},
		{name: "poison-middle", tbs: tbsPoisonMiddle, want: tbsNoPoison},
	}
	for _, test := range tests {
		in, _ := hex.DecodeString(test.tbs)
		got, err := RemoveCTPoison(in)
		if test.errstr != "" {
			if err == nil {
				t.Errorf("RemoveCTPoison(%s)=%s,nil; want error %q", test.name, hex.EncodeToString(got), test.errstr)
			} else if !strings.Contains(err.Error(), test.errstr) {
				t.Errorf("RemoveCTPoison(%s)=nil,%q; want error %q", test.name, err, test.errstr)
			}
			continue
		}
		want, _ := hex.DecodeString(test.want)
		if err != nil {
			t.Errorf("RemoveCTPoison(%s)=nil,%q; want %s,nil", test.name, err, test.want)
		} else if !bytes.Equal(got, want) {
			t.Errorf("RemoveCTPoison(%s)=%s,nil; want %s,nil", test.name, hex.EncodeToString(got), test.want)
		}
	}
}
