package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/d-smith/go-examples/awsreg"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: publish <topic arn> <message>")
		os.Exit(1)
	}

	svc := sns.New(session.New(), &aws.Config{Region: awsreg.RegionFromEnvOrDefault("us-east-1")})

	params := &sns.PublishInput{
		Message:  aws.String(os.Args[2]),
		Subject:  aws.String("a subject"),
		TopicArn: aws.String(os.Args[1]),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
