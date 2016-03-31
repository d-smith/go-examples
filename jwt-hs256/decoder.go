package main
import (
	"os"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"encoding/base64"
)


func checkArgs() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: decoder <secret> <token>")
	}
}

func decodeToken(secret, tokenString string) (*jwt.Token) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		key,err := base64.StdEncoding.DecodeString(secret)
		if err != nil {
			return nil,err
		}

		return key, nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return token

}

func printClaims(token *jwt.Token) {
	fmt.Println("Token claims...")
	for k,v := range token.Claims {
		fmt.Printf("%s -> %v\n",k,v)
	}
}

func main() {
	checkArgs();

	token := decodeToken(os.Args[1],os.Args[2])
	printClaims(token)
}
