package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
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

type KinesisEvent struct {
	Name string `json:"name"`
}

func putRecord(client *firehose.Firehose, streamName *string, text string) error {
	data := []byte(text)

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

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) {
	streamName := aws.String("TestFirehose")
	client := setClient()

	for _, record := range kinesisEvent.Records {
		kinesisRecord := record.Kinesis
		dataBytes := kinesisRecord.Data
		dataText := string(dataBytes)

		fmt.Printf("%s data: %s\n", record.EventName, dataText)
	}

	putRecord(client, streamName, "キネシスから入力")
}

func main() {
	lambda.Start(handler)
}
