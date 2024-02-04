package usecases

import (
	"fmt"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/pkg/utils"
)

type CreateAccountRequestDTO struct {
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Birthdate   string `json:"birthdate"`
}

type ClientDTO struct {
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Birthdate   string `json:"birthdate"`
}

type CreateAccountResponseDTO struct {
	Id      string    `json:"id"`
	Balance float64   `json:"balance"`
	Client  ClientDTO `json:"client"`
}

type CreateAccountUseCase struct{}

func (us *CreateAccountUseCase) Execute(accountDTO CreateAccountRequestDTO) (*CreateAccountResponseDTO, error) {
	birthdate, err := utils.ParseStringToTIme(accountDTO.Birthdate)
	if err != nil {
		return nil, err
	}
	client := entities.Client{
		Name:        accountDTO.Name,
		CPF:         accountDTO.Name,
		Email:       accountDTO.Email,
		PhoneNumber: accountDTO.PhoneNumber,
		Birthdate:   *birthdate,
	}
	fmt.Printf("%v\n", client)
	return &CreateAccountResponseDTO{
		Id:      "batata",
		Balance: 0.00,
		Client: ClientDTO{
			Name:        client.Name,
			CPF:         client.CPF,
			Email:       client.Email,
			PhoneNumber: client.PhoneNumber,
			Birthdate:   utils.FormatDate(*birthdate),
		},
	}, nil
}
