/*
Note - this example borrows heavily from github.com/dgrijalva/jwt-go
*/
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func extractPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	//For this sample we assume our key is PEM encoded
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("Key must be PEM encoded for this sample")
	}

	//Extract the private key. Note there's a couple ways this migbt be embedded in the key
	var parsedKey interface{}
	var err error
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errors.New("Unable to parse key from key bytes")
	}

	return pkey, nil

}

func extractPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	//For this sample we assume our key is PEM encoded
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("Key must be PEM encoded for this sample")
	}

	//Extract the public key. Note there's a couple ways it might be embedded in the key
	var parsedKey interface{}
	var err error
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, errors.New("Unable to parse key from key bytes")
	}

	return pkey, nil
}

func init() {
	signKeyBytes, err := ioutil.ReadFile("./testkey.rsa")
	fatal(err)

	signKey, err = extractPrivateKeyFromPEM(signKeyBytes)
	fatal(err)

	verifyKeyBytes, err := ioutil.ReadFile("./testkey.pub")
	fatal(err)

	verifyKey, err = extractPublicKeyFromPEM(verifyKeyBytes)
	fatal(err)
}

func signString(s string, key *rsa.PrivateKey) (string, error) {
	var hash crypto.Hash = crypto.SHA256
	if !hash.Available() {
		return "", errors.New("Hash not available")
	}

	hasher := hash.New()
	hasher.Write([]byte(s))

	signedBytes, err := rsa.SignPKCS1v15(rand.Reader, key, hash, hasher.Sum(nil))
	if err != nil {
		return "", err
	}

	return encodeSegment(signedBytes), nil

}

func verifyString(s, signature string, key *rsa.PublicKey) error {
	sig, err := decodeSegment(signature)
	if err != nil {
		return nil
	}

	var hash crypto.Hash = crypto.SHA256
	if !hash.Available() {
		return errors.New("Hash not available")
	}

	hasher := hash.New()
	hasher.Write([]byte(s))

	return rsa.VerifyPKCS1v15(key, hash, hasher.Sum(nil), sig)

}

//Taken directly from github.com/dgrijalva/jwt-go
func encodeSegment(seg []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
}

//Taken directly from github.com/dgrijalva/jwt-go
// Decode JWT specific base64url encoding with padding stripped
func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}

func main() {
	println("sign something")
	s := "This is something to sign."
	signed, err := signString(s, signKey)
	fatal(err)
	println(signed)

	println("verify something")
	err = verifyString(s, signed, verifyKey); if err != nil {
		println(err.Error())
	}

	println("signature verified")

	println("try to verify something that's been tampered with")
	err = verifyString("what's all this then?", signed, verifyKey); if err == nil {
		println("hmmmm... expected an error")
	} else {
		println(err.Error())
	}
}
