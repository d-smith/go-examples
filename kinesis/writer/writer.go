package main

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"os"
	"strconv"
	"time"
)

func writeLineToStream(svc *kinesis.Kinesis, streamName string, line string) {

	partitionKey := strconv.FormatInt(time.Now().UnixNano(), 10)

	params := &kinesis.PutRecordInput{
		Data:         []byte(line),
		PartitionKey: aws.String(partitionKey),
		StreamName:   aws.String(streamName),
	}
	_, err := svc.PutRecord(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println("error: " + err.Error())
		return
	}

}

func main() {

	svc := kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		writeLineToStream(svc, "test-stream", line)
	}
}
