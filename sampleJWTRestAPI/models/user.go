package models

import "time"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Tel       string `json:"tel"`
	Mail      string `json:"mail"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser() *User {
	return &User{}
}
