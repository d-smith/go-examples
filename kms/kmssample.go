package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"fmt"
)

func main() {
	encrypterRoleArn := os.Getenv("ENC_ROLE_ARN")
	descrypterRoleArn := os.Getenv("DEC_ROLE_ARN")

	if encrypterRoleArn == "" || descrypterRoleArn == "" {
		fmt.Println("Must provide ENC_ROLE_ARN and DEC_ROLE_ARN environment variables")
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())
	println(sess)
}
