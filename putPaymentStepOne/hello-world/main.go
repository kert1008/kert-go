package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/aws/aws-lambda-go/lambda"
)

type PaymentEvent struct {
	PaymentID     string      `json:"payment_id"`
	PaymentInfo   string      `json:"payment_info"`
	PaymentType   string      `json:"payment_type"`
	Amount        json.Number `json:"amount"`
	PaymentResult string      `json:"payment_result"`
}

func HandleRequest(ctx context.Context, event PaymentEvent) (PaymentEvent, error) {
	fmt.Printf("S1: PaymentID: [%v]\n", event.PaymentID)
	fmt.Printf("S1: PaymentInfo: [%v]\n", event.PaymentInfo)
	fmt.Printf("S1: PaymentType: [%v]\n", event.PaymentType)
	fmt.Printf("S1: Amount: [%v]\n", event.Amount)

	event.PaymentResult = judgePayment()
	return event, nil
}

func main() {
	lambda.Start(HandleRequest)
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
