package cryptoutil

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// Implementation of Java's SHA256WithRSA in Go
// Algorithm used are RSA-PKCS#1v1.5 and SHA256
func SHA256WithRSA(msg []byte, key *rsa.PublicKey) (res []byte, err error) {
	digest := sha256.Sum256(msg)

	res, err = rsa.EncryptPKCS1v15(rand.Reader, key, digest[:])
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt hash err: %+v", err)
	}

	return
}

func HMACSHA512(msg []byte, key []byte) []byte {
	return hmac.New(sha512.New, key).Sum(msg)
}

func LoadPublicKey(key []byte) (pk *rsa.PublicKey, err error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("failed to decode public key err: no public key found")
	}

	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key err: %+v", err)
	}

	pk = keyInterface.(*rsa.PublicKey)

	return
}
