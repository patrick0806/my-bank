package usecases

import (
	"errors"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/internal/respositories"
	"github.com/patrick0806/my-bank/pkg/entity"
	"github.com/patrick0806/my-bank/pkg/utils"
)

type CreateAccountRequestDTO struct {
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Birthdate   string `json:"birthdate"`
}

type CreateAccountResponseDTO struct {
	Id          string  `json:"id"`
	Balance     float64 `json:"balance"`
	Name        string  `json:"name"`
	CPF         string  `json:"cpf"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Birthdate   string  `json:"birthdate"`
}

type CreateAccountUseCase struct {
	AccountRepository respositories.AccountRepository
}

func NewCreateAccountUseCase(repository respositories.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: repository,
	}
}

func (us *CreateAccountUseCase) Execute(accountDTO CreateAccountRequestDTO) (*CreateAccountResponseDTO, error) {
	existAccount, err := us.AccountRepository.FindByCPF(accountDTO.CPF)
	if err != nil {
		return nil, err
	}

	if existAccount != nil {
		return nil, errors.New("already exists a account in my bank with this cpf")
	}

	birthdate, err := utils.ParseStringToTIme(accountDTO.Birthdate)
	if err != nil {
		return nil, err
	}

	hash, err := utils.HashPassword(accountDTO.Password)
	if err != nil {
		return nil, err
	}

	newAccount := &entities.Account{
		Id:          entity.NewID(),
		Balance:     0.00,
		Name:        accountDTO.Name,
		CPF:         accountDTO.CPF,
		Email:       accountDTO.Email,
		Password:    hash,
		PhoneNumber: accountDTO.PhoneNumber,
		Birthdate:   *birthdate,
	}

	err = us.AccountRepository.CreateAccount(newAccount)
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponseDTO{
		Id:          newAccount.Id.String(),
		Balance:     newAccount.Balance,
		Name:        newAccount.Name,
		CPF:         newAccount.CPF,
		Email:       newAccount.Email,
		PhoneNumber: newAccount.PhoneNumber,
		Birthdate:   utils.FormatDate(newAccount.Birthdate),
	}, nil
}
