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
	if len(os.Args) != 2 {
		fmt.Println("Usage: create-topic <topic name>")
		return
	}
	svc := sns.New(session.New(), &aws.Config{Region: awsreg.RegionFromEnvOrDefault("us-east-1")})

	params := &sns.CreateTopicInput{
		Name: aws.String(os.Args[1]),
	}
	resp, err := svc.CreateTopic(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Topic arn:", *resp.TopicArn)
}
