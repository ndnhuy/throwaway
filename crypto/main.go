// main.go
package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"
)

var pemPrivateKey = `
`

var pemPub = ``

func main() {
	body := `{
		"additionalInfo": {
			"accessToken": "",
			"referenceNo": ""
		},
		"originalParnerReferenceNo": ""
	}`
	p := payload(body)
	fmt.Println("[Body]\n" + minify(body))
	fmt.Println("\n[Payload]\n" + p)
	fmt.Printf("\n[Signature]\n%v", encodeBase64(sign(p)))

	block, _ := pem.Decode([]byte(pemPub))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("error getting public key")
	}

	fmt.Printf("\n\nVerify signature with public key: %v", verify(p, sign(p), pub.(*rsa.PublicKey)))
}

func sign(payload string) []byte {
	block, _ := pem.Decode([]byte(pemPrivateKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
		panic("error getting private key")
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashSHA256(payload))
	if err != nil {
		panic(err)
	}

	return signature
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func verify(msg string, sig []byte, pub *rsa.PublicKey) bool {
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashSHA256(msg), sig)
	if err != nil {
		panic("error verifying signature")
	}
	return true
}

func hashSHA256(msg string) []byte {
	h := sha256.New()
	h.Write([]byte(msg))
	d := h.Sum(nil)
	return d
}

// <HTTP METHOD> + ”:” + <RELATIVE PATH URL> + “:“ + LowerCase(HexEncode(SHA-256(Minify(<HTTP BODY>)))) + “:“ + <X-TIMESTAMP>
func payload(body string) string {
	// clientId := "2022070816171583137012"
	timestamp := "2023-04-17T11:01:43+07:00"
	httpMethod := "POST"
	relativePath := "/v1.0/transaction-history-detail.htm"
	p := httpMethod + ":" + relativePath + ":" + lowercase(hex.EncodeToString(hashSHA256(minify(body)))) + ":" + timestamp
	return p
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

func minify(jsonBody string) string {
	// compact json
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(jsonBody))
	if err != nil {
		panic(err)
	}
	return buf.String()
}
