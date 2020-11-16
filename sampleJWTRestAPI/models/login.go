package models

import "time"

type Login struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLogin() *Login {
	return &Login{}
}
