package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: subscribe <topic arn> <queue arn>")
	}

	svc := sns.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	params := &sns.SubscribeInput{
		Protocol: aws.String("sqs"),      // Required
		TopicArn: aws.String(os.Args[1]), // Required
		Endpoint: aws.String(os.Args[2]),
	}
	resp, err := svc.Subscribe(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}
