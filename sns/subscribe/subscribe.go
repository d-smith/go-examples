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
		return
	}

	svc := sns.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	params := &sns.SubscribeInput{
		Protocol: aws.String("sqs"),
		TopicArn: aws.String(os.Args[1]),
		Endpoint: aws.String(os.Args[2]),
	}
	resp, err := svc.Subscribe(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

}
