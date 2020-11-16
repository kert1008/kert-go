package controller

import (
	"fmt"
	"hello-world/models"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func GetPayment(paymentID string) (error, models.Payment) {
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

	table := db.Table("payment")

	var readResult models.Payment
	var err error

	for i := 0; i < 10; i++ {
		fmt.Printf("i=[%v]\n", i)
		err = table.Get("payment_id", paymentID).One(&readResult)
		if err == nil {
			break
		}
		fmt.Printf("GetPaymentError = [%v]\n", err)
		time.Sleep(time.Millisecond * 500)
	}

	return err, readResult
}
