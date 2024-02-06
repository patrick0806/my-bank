package usecases_test

import (
	"testing"
	"time"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/entity"
	"github.com/stretchr/testify/assert"
)

var mockAccount = entities.Account{
	Id:          entity.NewID(),
	Balance:     0.00,
	Name:        "Teste Account",
	Password:    "$2a$14$CVoI0pwtpL6Iz0VJvY.NWOdpHQ.Q1JckrC49mpUAqssCG11cnAvb.",
	CPF:         "00000000000",
	Email:       "test@email.com",
	PhoneNumber: "99999990000",
	Birthdate:   time.Now(),
}

func TestShouldSignInWithSuccesss(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	signInUseCase := usecases.SignInUseCase{
		AccountRepository: mockAccountRepository,
	}

	mockAccountRepository.On("FindByCPF", "00000000000").Return(&mockAccount, nil)

	requestDTO := usecases.SignRequestDTO{
		CPF:      "00000000000",
		Password: "batata",
	}
	accessToken, err := signInUseCase.Execute(requestDTO)
	assert.Nil(t, err)
	assert.NotNil(t, accessToken)
}

func TestShouldFailSignInWithWrongPassword(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	signInUseCase := usecases.SignInUseCase{
		AccountRepository: mockAccountRepository,
	}

	mockAccountRepository.On("FindByCPF", "00000000000").Return(&mockAccount, nil)

	requestDTO := usecases.SignRequestDTO{
		CPF:      "00000000000",
		Password: "potato",
	}
	accessToken, err := signInUseCase.Execute(requestDTO)
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, usecases.ErrInvalidLoginOrPassword)
}

func TestShouldFailWithNotFoundCPF(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	signInUseCase := usecases.SignInUseCase{
		AccountRepository: mockAccountRepository,
	}

	mockAccountRepository.On("FindByCPF", "00000000000").Return(nil, nil)

	requestDTO := usecases.SignRequestDTO{
		CPF:      "00000000000",
		Password: "batata",
	}
	accessToken, err := signInUseCase.Execute(requestDTO)
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, usecases.ErrInvalidLoginOrPassword)
}
