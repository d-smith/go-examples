package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("gen us some keys")

	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	fatal(err)

	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	fmt.Println(string(pemdata))

	pubkey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	fatal(err)

	pemdata = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey,
		},
	)

	fmt.Println(string(pemdata))
}
