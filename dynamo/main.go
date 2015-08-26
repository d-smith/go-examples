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

func (d *Developer) Create(client *dynamodb.DynamoDB) error {
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

func (d *Developer) Update(first, last string, client *dynamodb.DynamoDB) error {
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("Developer"),
		Key: map[string]*dynamodb.AttributeValue{
			"EMail": {S: aws.String(d.EMail)},
		},
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"FirstName": {
				Action: aws.String(dynamodb.AttributeActionPut),
				Value: &dynamodb.AttributeValue{
					S: aws.String(first),
				},
			},
			"LastName": {
				Action: aws.String(dynamodb.AttributeActionPut),
				Value: &dynamodb.AttributeValue{
					S: aws.String(last),
				},
			},
		},
	}

	_, err := client.UpdateItem(params)

	return err
}

func Get(email string, client *dynamodb.DynamoDB) (*Developer, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Developer"),
		Key: map[string]*dynamodb.AttributeValue{
			"EMail": {S: aws.String(email)},
		},
	}

	out, err := client.GetItem(params)
	if err != nil {
		return nil, err
	}

	return &Developer{
		EMail:     *out.Item["EMail"].S,
		FirstName: *out.Item["FirstName"].S,
		LastName:  *out.Item["LastName"].S,
	}, nil
}

func Delete(email string, client *dynamodb.DynamoDB) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Developer"),
		Key: map[string]*dynamodb.AttributeValue{
			"EMail": {S: aws.String(email)},
		},
	}

	_, err := client.DeleteItem(params)
	return err
}

func main() {
	client := dynamodb.New(&aws.Config{Region: aws.String("us-east-1")})
	devEmail := "dev@dev.com"

	dev := Developer{
		EMail:     devEmail,
		FirstName: "Joe",
		LastName:  "Dev",
	}

	err := dev.Create(client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	fmt.Println("dev stored")

	retdev, err := Get(devEmail, client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	fmt.Println(retdev)

	err = dev.Update("updated first", "updated last", client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	retdev, err = Get(devEmail, client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	fmt.Println(retdev)

	err = Delete(devEmail, client)
	if err != nil {
		fmt.Println("Dang it: ", err.Error())
		return
	}

	fmt.Println("deleted")
}
