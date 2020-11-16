package models

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DynamoUser struct {
	ID      string `dynamo:"id"`
	Name    string `dynamo:"name"`
	Address string `dynamo:"address"`
}
