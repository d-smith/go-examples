package main

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
	"time"
)

const (
	streamName = "test-stream"
)

func getShardId(svc *kinesis.Kinesis) (string, error) {
	params := &kinesis.DescribeStreamInput{
		StreamName:            aws.String(streamName), // Required
		ExclusiveStartShardId: aws.String("ShardId"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.DescribeStream(params)
	if err != nil {
		return "", err
	}

	if len(resp.StreamDescription.Shards) != 1 {
		return "", errors.New("This sample is prepared to deal with only one shard")
	}

	shardId := resp.StreamDescription.Shards[0].ShardId

	return *shardId, nil
}

func getShardItor(svc *kinesis.Kinesis, sid string) (*string, error) {
	params := &kinesis.GetShardIteratorInput{
		ShardId:           aws.String(sid),            // Required
		ShardIteratorType: aws.String("TRIM_HORIZON"), // Required
		StreamName:        aws.String(streamName),     // Required
	}

	itor, err := svc.GetShardIterator(params)
	if err != nil {
		return nil, err
	}

	return itor.ShardIterator, nil
}

func readRecordsFromStream(svc *kinesis.Kinesis, shardItor *string) (*string, error) {

	params := &kinesis.GetRecordsInput{
		ShardIterator: shardItor, // Required
		Limit:         aws.Int64(1),
	}
	resp, err := svc.GetRecords(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return nil, err
	}

	records := resp.Records
	for _, r := range records {
		log.Printf(">>>> %s\n", string(r.Data))
	}

	return resp.NextShardIterator, nil
}

func main() {

	svc := kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	sid, err := getShardId(svc)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Working with shard", sid)
	shardItor, err := getShardItor(svc, sid)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		log.Println("read some records...")
		shardItor, err = readRecordsFromStream(svc, shardItor)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("that was tiring... time for nap")
		time.Sleep(10 * time.Second)
	}
}
