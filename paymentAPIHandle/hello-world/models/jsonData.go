package models

import "encoding/json"

type PaymentRequest struct {
	PaymentInfo string      `json:"payment_info"`
	PaymentType string      `json:"payment_type"`
	Amount      json.Number `json:"amount"`
}

type PaymentResponse struct {
	PaymentID     string      `json:"payment_id"`
	PaymentInfo   string      `json:"payment_info"`
	PaymentType   string      `json:"payment_type"`
	Amount        json.Number `json:"amount"`
	PaymentResult string      `json:"payment_result"`
}
