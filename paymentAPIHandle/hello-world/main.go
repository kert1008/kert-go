package main

import (
	"bytes"
	"context"
	"fmt"
	"hello-world/controller"
	"hello-world/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.GET("/", createPayment())

	ginLambda = ginadapter.New(r)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handler)
}

func createPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		tmpBody := string(buf[0:n])
		fmt.Printf("PaymentRequest body = [%s]\n", tmpBody)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(tmpBody)))
		paymentRequest := models.PaymentRequest{}
		c.BindJSON(&paymentRequest)

		paymentID := generatePaymentID()

		if controller.PutPaymentEvent(&paymentRequest, paymentID) == nil {
			err, payment := getPaymentResult(paymentID)
			if err != nil {
				c.JSON(http.StatusOK, "Payment failed.")
			} else {
				c.JSON(http.StatusOK, payment)
			}

		} else {
			c.JSON(http.StatusOK, "Payment failed.")
		}
	}
}

func getPaymentResult(paymentID string) (error, models.Payment) {

	fmt.Printf("paymentID to search: paymentID=[%s]\n", paymentID)
	err, payment := controller.GetPayment(paymentID)

	return err, payment
}

func generatePaymentID() string {
	const MilliFormat = "20060102150405000"
	paymentID := "KTPAY" + time.Now().Format(MilliFormat)
	fmt.Println(paymentID)
	return paymentID
}
