package usecases

import (
	"errors"

	"github.com/patrick0806/my-bank/internal/respositories"
	"github.com/patrick0806/my-bank/pkg/utils"
)

type SignRequestDTO struct {
	CPF      string `json:"cpf"`
	Password string `json:"password"`
}

type SignResponseDTO struct {
	AccessToken string `json:"accessToken"`
}

type SignInUseCase struct {
	AccountRepository respositories.AccountRepository
}

func NewSignInUseCase(respository respositories.AccountRepository) *SignInUseCase {
	return &SignInUseCase{
		AccountRepository: respository,
	}
}

var ErrInvalidLoginOrPassword = errors.New("invalid username or password")

func (us *SignInUseCase) Execute(loginData SignRequestDTO) (*SignResponseDTO, error) {
	account, err := us.AccountRepository.FindByCPF(loginData.CPF)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, ErrInvalidLoginOrPassword
	}

	isValidPassword := utils.CheckPasswordHash(loginData.Password, account.Password)
	if !isValidPassword {
		return nil, ErrInvalidLoginOrPassword
	}

	stringToken, err := utils.GenerateJWT(account.Id.String(), account.Name, account.Email)
	if err != nil {
		return nil, err
	}

	return &SignResponseDTO{
		AccessToken: stringToken,
	}, nil
}
