package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type PaymentEvent struct {
	PaymentID     string `json:"payment_id"`
	PaymentInfo   string `json:"payment_info"`
	PaymentType   string `json:"payment_type"`
	Amount        int    `json:"amount"`
	PaymentResult string `json:"payment_result"`
}

type PaymentDynamo struct {
	PaymentID     string `dynamo:"payment_id"`
	PaymentInfo   string `dynamo:"payment_info"`
	PaymentType   string `dynamo:"payment_type"`
	Amount        int    `dynamo:"amount"`
	PaymentResult string `dynamo:"payment_result"`
}

func HandleRequest(ctx context.Context, event PaymentEvent) error {
	fmt.Printf("S2: PaymentID: [%v]\n", event.PaymentID)
	fmt.Printf("S2: PaymentInfo: [%v]\n", event.PaymentInfo)
	fmt.Printf("S2: PaymentType: [%v]\n", event.PaymentType)
	fmt.Printf("S2: Amount: [%v]\n", event.Amount)
	fmt.Printf("S2: PaymentResult: [%v]\n", event.PaymentResult)

	err := addPayment(event)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
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
	dPayment := PaymentDynamo{PaymentID: payment.PaymentID, PaymentInfo: payment.PaymentInfo, PaymentType: payment.PaymentType, Amount: payment.Amount, PaymentResult: payment.PaymentResult}
	err := table.Put(dPayment).Run()

	return err
}
