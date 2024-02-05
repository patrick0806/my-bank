package entities

import (
	"time"

	"github.com/patrick0806/my-bank/pkg/entity"
)

type Account struct {
	Id          entity.ID
	Balance     float64
	Name        string
	CPF         string
	Email       string
	PhoneNumber string
	Birthdate   time.Time
}
