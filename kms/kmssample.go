package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"crypto/aes"

	"io"
	"crypto/cipher"
	"crypto/rand"
	"errors"
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

func basicDemo(svc *kms.KMS) {
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

//Encrypt and Decrypt from cryptopasta commit bc3a108a5776376aa811eea34b93383837994340
//used via the CC0 license.
func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

func genKeyDemo(svc *kms.KMS) {
	params := &kms.GenerateDataKeyInput{
		KeyId: aws.String("alias/keyalias"), // Required
		KeySpec:       aws.String("AES_256"),
	}
	resp, err := svc.GenerateDataKey(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

	key := [32]byte{}

	copy(key[:],key[0:32])

	encrypted,err := Encrypt([]byte("a test"),&key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	decypted, err := Decrypt(encrypted, &key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Decrypted :", string(decypted))
}

func main() {

	sess := session.Must(session.NewSession())

	svc := kms.New(sess)
	//basicDemo(svc)

	genKeyDemo(svc)
}
