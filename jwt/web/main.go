package main

import (
	"net/http"
	"flag"
	"fmt"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"log"

)

const (
	privateKeyPath = "../keys/app.rsa" 	//openssl genrsa -out app.rsa 1024
	publicKeyPath = "../keys/app.rsa.pub" //openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
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


}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func extractTokenFromHeaderValue(tokenString string)(*jwt.Token,error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey,nil
	})

	return token,err
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	//Grab the Xt-ApiKey header
	apiKey := r.Header.Get("Xt-ApiKey")
	println("api key: ", apiKey)
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Extract the token from the token string
	token, err := extractTokenFromHeaderValue(apiKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//Is the token valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Look at some token info
	for k,v := range token.Claims {
		log.Println("Claim ", k, " is ", v)
	}

	w.Write([]byte("Request looks legit - here are the goods"))
}

func main() {
	var port = flag.Int("port", -1, "Port to listen on")
	flag.Parse()
	if *port == -1 {
		fmt.Println("Must specify a -port argument")
		return
	}

	http.Handle("/goods", http.HandlerFunc(handleRequest))
	http.ListenAndServe(fmt.Sprintf(":%d",*port), nil)

}
