package main

//Note - this code borrows heavily from https://gist.github.com/cryptix/45c33ecf0ae54828e63b

import (
	"crypto/rsa"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

const (
	privateKeyPath      = "../keys/app.rsa"          //openssl genrsa -out app.rsa 1024
	publicKeyPath       = "../keys/app.rsa.pub"      //openssl rsa -in app.rsa -pubout > app.rsa.pub
	otherKeyPath        = "../keys/otherkey.rsa.pub" //openssl genrsa -out otherkey.rsa 1024
	otherPrivateKeyPath = "../keys/otherkey.rsa"     //openssl rsa -in otherkey.rsa -pubout > otherkey.rsa.pub
)

var (
	verifyKey      *rsa.PublicKey
	signKey        *rsa.PrivateKey
	otherVerifyKey *rsa.PublicKey
	otherSignKey   *rsa.PrivateKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

	otherKeyBytes, err := ioutil.ReadFile(otherKeyPath)
	fatal(err)

	otherVerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(otherKeyBytes)
	fatal(err)

	otherSignBytes, err := ioutil.ReadFile(otherPrivateKeyPath)
	fatal(err)

	otherSignKey, err = jwt.ParseRSAPrivateKeyFromPEM(otherSignBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["AccessToken"] = "level1"
	t.Claims["ApplicationName"] = "foo app"
	t.Claims["App Registry ID"] = "AP0001"

	tokenString, err := t.SignedString(signKey)
	fatal(err)

	log.Println("token string - sign key: ", tokenString)

	otherTokenString, err := t.SignedString(otherSignKey)
	fatal(err)
	log.Println("token string - other: ", otherTokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	fatal(err)

	log.Println("Token valid: ", token.Valid)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return otherVerifyKey, nil
	})

	if err != nil {
		log.Fatal("expected error when validating token with other key")
	}

	log.Println("As expected, error returned when validating token with other key: ", err.Error())

}
