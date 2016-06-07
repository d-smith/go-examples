package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"github.com/d-smith/go-examples/awsreg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: dequeue <queue url>")
		os.Exit(1)
	}

	svc := sqs.New(session.New(), &aws.Config{Region: awsreg.RegionFromEnvOrDefault("us-east-1")})

	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(os.Args[1]),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout: aws.Int64(1),
		WaitTimeSeconds:   aws.Int64(1),
	}
	resp, err := svc.ReceiveMessage(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.Messages)
	if len(resp.Messages) == 0 {
		return
	}

	delParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Args[1]),
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	}
	delResp, err := svc.DeleteMessage(delParams)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(delResp)
}
