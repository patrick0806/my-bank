package entities

import "time"

type Client struct {
	Id          string
	Name        string
	CPF         string
	Email       string
	PhoneNumber string
	Birthdate   time.Time
}
