package controller

import (
	"fmt"
	"hello-world/models"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
)

func PutPaymentEvent(req *models.PaymentRequest, paymentID string) error {

	mySession := session.Must(session.NewSession())
	svc := eventbridge.New(mySession)

	putEventsRequestEntry := eventbridge.PutEventsRequestEntry{}
	putEventsRequestEntry.SetSource("PaymentAPIHandleRequest")
	putEventsRequestEntry.SetDetailType("CreatePaymentInfo")

	amountInt, _ := req.Amount.Int64()
	amountStr := strconv.FormatInt(amountInt, 10)

	reqDetail := "{ \"payment_id\": \"" + paymentID +
		"\", \"payment_info\": \"" + req.PaymentInfo +
		"\", \"payment_type\": \"" + req.PaymentType +
		"\", \"amount\": " + amountStr + " }"

	fmt.Printf("PutEvents detail = [%s]\n", reqDetail)

	putEventsRequestEntry.SetDetail(reqDetail)

	putEventsRequestEntryList := []*eventbridge.PutEventsRequestEntry{&putEventsRequestEntry}

	putEventsInput := eventbridge.PutEventsInput{}
	putEventsInput.SetEntries(putEventsRequestEntryList)

	result, err := svc.PutEvents(&putEventsInput)
	fmt.Printf("PutEvents result = [%v]\n", result)
	fmt.Printf("PutEvents error = [%v]\n", err)
	return err
}
