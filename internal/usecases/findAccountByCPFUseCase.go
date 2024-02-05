package usecases

import (
	"github.com/patrick0806/my-bank/internal/respositories"
	"github.com/patrick0806/my-bank/pkg/utils"
)

type FindAccountByCPFUseCase struct {
	AccountRepository respositories.AccountRepository
}

type FindAccountByCPFResponseDTO struct {
	Id          string  `json:"id"`
	Balance     float64 `json:"balance"`
	Name        string  `json:"name"`
	CPF         string  `json:"cpf"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Birthdate   string  `json:"birthdate"`
}

func NewFindAccountByCPFUseCase(repository *respositories.AccountRepository) *FindAccountByCPFUseCase {
	return &FindAccountByCPFUseCase{
		AccountRepository: *repository,
	}
}

func (us *FindAccountByCPFUseCase) Execute(cpf string) (*FindAccountByCPFResponseDTO, error) {
	account, err := us.AccountRepository.FindByCPF(cpf)
	if err != nil {
		return nil, err
	}
	return &FindAccountByCPFResponseDTO{
		Id:          account.Id.String(),
		Name:        account.Name,
		CPF:         account.CPF,
		Balance:     account.Balance,
		Email:       account.Email,
		PhoneNumber: account.PhoneNumber,
		Birthdate:   utils.FormatDate(account.Birthdate),
	}, nil
}
