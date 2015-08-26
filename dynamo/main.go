package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Developer struct {
	EMail     string
	FirstName string
	LastName  string
}

func (d *Developer) Store(client *dynamodb.DynamoDB) error {
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Developer"),
		Item: map[string]*dynamodb.AttributeValue{
			"EMail":     {S: aws.String(d.EMail)},
			"FirstName": {S: aws.String(d.FirstName)},
			"LastName":  {S: aws.String(d.LastName)},
		},
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"EMail": {Exists: aws.Bool(false)},
		},
	}
	_, err := client.PutItem(params)

	return err
}

func main() {
	client := dynamodb.New(&aws.Config{Region: aws.String("us-east-1")})

	dev := Developer{
		EMail:     "dev@dev.com",
		FirstName: "Joe",
		LastName:  "Dev",
	}

	err := dev.Store(client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	fmt.Println("dev stored")
}
