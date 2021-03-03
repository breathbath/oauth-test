package enc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	defaultLength = 2048
)

var DefaultKeyProvider = func() (key *rsa.PrivateKey, err error) {
	key, err = rsa.GenerateKey(rand.Reader, defaultLength)
	if err != nil {
		return nil, err
	}

	fmt.Println(ExportRsaPrivateKeyAsPemStr(key))
	pubKey, err := ExportRsaPublicKeyAsPemStr(&key.PublicKey)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(pubKey)

	return
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}
