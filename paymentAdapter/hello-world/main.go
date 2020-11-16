package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type PaymentEvent struct {
	PaymentID   string `json:"payment_id"`
	PaymentInfo string `json:"payment_info"`
	PaymentType string `json:"payment_type"`
	Amount      int    `json:"amount"`
}

type PaymentDynamo struct {
	PaymentID     string `dynamo:"payment_id"`
	PaymentInfo   string `dynamo:"payment_info"`
	PaymentType   string `dynamo:"payment_type"`
	Amount        int    `dynamo:"amount"`
	PaymentResult string `dynamo:"payment_result"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		payment := PaymentEvent{}
		err := json.Unmarshal([]byte(message.Body), &payment)
		if err != nil {
			return err
		}

		err = addPayment(payment)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

func addPayment(payment PaymentEvent) error {
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

	fmt.Printf("Payment = [%v]\n", payment)

	table := db.Table("payment")
	dPayment := PaymentDynamo{PaymentID: payment.PaymentID, PaymentInfo: payment.PaymentInfo, PaymentType: payment.PaymentType, Amount: payment.Amount, PaymentResult: judgePayment()}
	err := table.Put(dPayment).Run()

	return err
}

func judgePayment() string {
	num := rand.Intn(100)
	fmt.Printf("num = [%v]\n", num)

	if num%2 == 0 {
		fmt.Printf("Payment succeeded.\n")
		return "succeed"
	}
	fmt.Printf("Payment failed.\n")
	return "failed"
}
