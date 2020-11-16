package models

type User struct {
	ID      string `dynamo:"id"`
	Name    string `dynamo:"name"`
	Address string `dynamo:"address"`
}
