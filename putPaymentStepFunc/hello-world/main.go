package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

type PaymentRequest struct {
	PaymentID   string      `json:"payment_id"`
	PaymentInfo string      `json:"payment_info"`
	PaymentType string      `json:"payment_type"`
	Amount      json.Number `json:"amount"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		newPayment := PaymentRequest{}
		err := json.Unmarshal([]byte(message.Body), &newPayment)
		if err != nil {
			return err
		}

		startSFExecution(&newPayment)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

func startSFExecution(payment *PaymentRequest) error {
	mySession := session.Must(session.NewSession())
	svc := sfn.New(mySession)

	amountInt, _ := payment.Amount.Int64()
	amountStr := strconv.FormatInt(amountInt, 10)

	sfInput := sfn.StartExecutionInput{}

	input := "{ \"payment_id\": \"" + payment.PaymentID +
		"\", \"payment_info\": \"" + payment.PaymentInfo +
		"\", \"payment_type\": \"" + payment.PaymentType +
		"\", \"amount\": " + amountStr + " }"

	sfInput.SetInput(input)
	sfInput.SetName("MyStateMachine" + payment.PaymentID)
	sfInput.SetStateMachineArn("##Your StateMachine's ARN##")

	result, err := svc.StartExecution(&sfInput)
	fmt.Printf("StartExecution result = [%v]\n", result)
	fmt.Printf("StartExecution error = [%v]\n", err)
	return err
}
