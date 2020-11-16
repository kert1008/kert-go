package controller

import (
	"bytes"
	"fmt"
	"hello-world/models"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

func ReadUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		tmpBody := string(buf[0:n])
		fmt.Printf("request body = [%s]\n", tmpBody)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(tmpBody)))
		input := models.User{}
		c.BindJSON(&input)
		fmt.Printf("inputUser: [%v]\n", input)

		result := addDynamodb(input)
		c.JSON(http.StatusOK, result)
	}
}

func addDynamodb(user models.User) string {
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
	dUser := models.DynamoUser{ID: user.ID, Name: user.Name, Address: user.Address}
	err := table.Put(dUser).Run()

	if err != nil {
		return "false"
	}
	return "true"
}
