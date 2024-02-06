package usecases_test

import (
	"testing"

	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldCreateAccount(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	createAccountUseCase := usecases.CreateAccountUseCase{
		AccountRepository: mockAccountRepository,
	}

	mockAccountRepository.On("FindByCPF", "00000000000").Return(nil, nil)
	mockAccountRepository.On("CreateAccount", mock.AnythingOfType("*entities.Account")).Return(nil)

	requestDTO := usecases.CreateAccountRequestDTO{
		Name:        "Teste User",
		CPF:         "00000000000",
		Email:       "test@test.com",
		Password:    "batata",
		PhoneNumber: "00000000000",
		Birthdate:   "2022-02-01",
	}

	createdAccount, err := createAccountUseCase.Execute(requestDTO)

	assert.Nil(t, err)
	assert.NotNil(t, createdAccount)
	assert.NotNil(t, createdAccount.Id)
	assert.Equal(t, 0.00, createdAccount.Balance)
	assert.Equal(t, requestDTO.Name, createdAccount.Name)
	assert.Equal(t, requestDTO.Email, createdAccount.Email)
	assert.Equal(t, requestDTO.CPF, createdAccount.CPF)
	assert.Equal(t, requestDTO.PhoneNumber, createdAccount.PhoneNumber)
	assert.Equal(t, requestDTO.Birthdate, createdAccount.Birthdate)
}
