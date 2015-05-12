package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "This is a test of base64 encoding, which is pretty easy"
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("Encoded: ", encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Print("Decoded: ", string(decoded))
	}
}
