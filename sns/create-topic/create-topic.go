package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	svc := sns.New(session.New(),&aws.Config{Region: aws.String("us-east-1")})

	params := &sns.CreateTopicInput{
		Name: aws.String("my-topic"),
	}
	resp, err := svc.CreateTopic(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Topic arn:", *resp.TopicArn)
}
