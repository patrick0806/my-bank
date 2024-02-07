package entities

import "github.com/patrick0806/my-bank/pkg/entity"

type Transaction struct {
	Id             entity.ID
	Value          float64
	AccountOrigin  string
	AccountDestiny string
}
