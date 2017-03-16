package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {

	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.EncryptInput{
		KeyId:     aws.String("alias/keyalias"),
		Plaintext: []byte("PAYLOAD"),
	}
	resp, err := svc.Encrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
