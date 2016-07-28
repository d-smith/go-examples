package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/d-smith/go-examples/awsreg"
	"os"
	"text/template"
)

var policyTemplate = `
{
  "Version": "2012-10-17",
  "Id": "{{.QueueArn}}/SQSDefaultPolicy",
  "Statement": [
    {
      "Sid": "Sid1462955527990",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "SQS:SendMessage",
      "Resource": "{{.QueueArn}}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "{{.TopicArn}}"
        }
      }
    }
  ]
}`

type policyCtx struct {
	QueueArn string
	TopicArn string
}

func createPolicy(queueArn, topicArn string) (string, error) {
	tmpl, err := template.New("policyTemplate").Parse(policyTemplate)
	if err != nil {
		return "", nil
	}

	ctx := policyCtx{
		QueueArn: queueArn,
		TopicArn: topicArn,
	}

	b := new(bytes.Buffer)

	err = tmpl.Execute(b, ctx)
	if err != nil {
		return "", nil
	}

	return b.String(), nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: pubperm <queue url> <queue arn> <topic arn>")
		return
	}

	policy, err := createPolicy(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(policy)

	svc := sqs.New(session.New(), &aws.Config{Region: awsreg.RegionFromEnvOrDefault("us-east-1")})

	params := &sqs.SetQueueAttributesInput{
		Attributes: map[string]*string{
			"Policy": aws.String(policy),
		},
		QueueUrl: aws.String(os.Args[1]), // Required
	}
	resp, err := svc.SetQueueAttributes(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
