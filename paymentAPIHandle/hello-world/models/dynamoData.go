package models

type Payment struct {
	PaymentID     string `dynamo:"payment_id"`
	PaymentInfo   string `dynamo:"payment_info"`
	PaymentType   string `dynamo:"payment_type"`
	Amount        int    `dynamo:"amount"`
	PaymentResult string `dynamo:"payment_result"`
}
