package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"success"`
}

func putRecord(client *firehose.Firehose, streamName *string) error {
	data := []byte("firehoseTest!")

	_, err := client.PutRecord(&firehose.PutRecordInput{
		DeliveryStreamName: streamName,
		Record:             &firehose.Record{Data: data},
	})
	if err != nil {
		return err
	}

	return nil
}

func setClient() *firehose.Firehose {
	return firehose.New(session.New(), &aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	})
}

func hello(event MyEvent) (MyResponse, error) {
	streamName := aws.String("TestFirehose")
	client := setClient()
	err := putRecord(client, streamName)
	if err != nil {
		return MyResponse{Message: fmt.Sprintf("ERROR %v", err)}, nil
	}
	return MyResponse{Message: fmt.Sprintf("put OK")}, nil
}

func main() {
	lambda.Start(hello)
}
