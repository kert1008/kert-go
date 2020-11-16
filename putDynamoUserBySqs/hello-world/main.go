package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DynamoUser struct {
	ID      string `dynamo:"id"`
	Name    string `dynamo:"name"`
	Address string `dynamo:"address"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		newUser := User{}
		err := json.Unmarshal([]byte(message.Body), &newUser)
		if err != nil {
			return err
		}

		err = addDynamodb(newUser)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

func addDynamodb(user User) error {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("user")
	dUser := DynamoUser{ID: user.ID, Name: user.Name, Address: user.Address}
	err := table.Put(dUser).Run()

	return err
}
