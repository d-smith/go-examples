package main


//Note - this code borrows heavily from https://gist.github.com/cryptix/45c33ecf0ae54828e63b

import (
	jwt "github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"log"
	"io/ioutil"
	"fmt"
)

const (
	privateKeyPath = "./app.rsa" 	//openssl genrsa -out app.rsa 1024
	publicKeyPath = "./app.rsa.pub" //openssl rsa -in app.rsa -pubout > app.rsa.pub
	otherKeyPath = "./otherkey.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
	otherKey  *rsa.PublicKey
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

	otherKey, err = jwt.ParseRSAPublicKeyFromPEM(otherKeyBytes)
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

	tokenString,err := t.SignedString(signKey)
	fatal(err)

	log.Println("token string: ", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey,nil
	})

	fatal(err)

	log.Println("Token valid: ", token.Valid)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return otherKey,nil
	})	

	if err != nil {
		log.Fatal("expected error when validating token with other key")
	}

	log.Println("As expected, error returned when validating token with other key: ", err.Error())

}
