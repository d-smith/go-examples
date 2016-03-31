package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
)

func GenerateID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func GenerateClientSecret() (string, error) {
	randbuf := make([]byte, 32)

	_, err := rand.Read(randbuf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(randbuf), nil
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	clientId, err := GenerateID()
	fatal(err)
	clientSecret, err := GenerateClientSecret()
	fatal(err)

	fmt.Printf("client id: %s, secret: %s\n", clientId, clientSecret)
}
