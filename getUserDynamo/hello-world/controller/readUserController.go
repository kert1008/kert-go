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

type Query struct {
	ID string `json:"id"`
}

func ReadUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		tmpBody := string(buf[0:n])
		fmt.Printf("request body = [%s]\n", tmpBody)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(tmpBody)))
		inputQuery := Query{}
		c.BindJSON(&inputQuery)
		fmt.Printf("Get id from body: id=[%s]\n", inputQuery.ID)

		user := readDynamodb(inputQuery.ID)
		c.JSON(http.StatusOK, user)
	}
}

func readDynamodb(id string) models.User {
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

	var readResult models.User
	if len(id) == 0 {
		id = "1"
	}
	fmt.Printf("id to search: id=[%s]\n", id)

	err := table.Get("id", id).One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
	} else {
		fmt.Printf("result item = [%v]\n", readResult)
	}

	return readResult
}
