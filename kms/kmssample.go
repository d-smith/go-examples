package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

func encryptStuff(svc *kms.KMS, plainText string) ([]byte,error) {
	params := &kms.EncryptInput{
		KeyId:     aws.String("alias/keyalias"),
		Plaintext: []byte(plainText),
	}

	resp, err:= svc.Encrypt(params)
	if err != nil {
		return nil, err
	}

	return resp.CiphertextBlob, nil
}

func decryptStuff(svc *kms.KMS, encrypted []byte) (string, error) {
	params := &kms.DecryptInput{
		CiphertextBlob: encrypted, // Required
	}
	resp, err := svc.Decrypt(params)

	if err != nil {
		return "", err
	}

	return string(resp.Plaintext), nil
}

func main() {

	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	enc, err := encryptStuff(svc, "My secret stuff")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(enc)

	dec, err := decryptStuff(svc,enc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(dec)

}
