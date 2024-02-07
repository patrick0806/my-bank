package usecases_test

import (
	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) FindByCPF(cpf string) (*entities.Account, error) {
	args := m.Called(cpf)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Account), args.Error(1)
}

func (m *MockAccountRepository) CreateAccount(accountData *entities.Account) error {
	args := m.Called(accountData)
	return args.Error(0)
}

func (m *MockAccountRepository) UpdateBalance(id string, balance float64) error {
	return nil
}
func (m *MockAccountRepository) FindById(id string) (*entities.Account, error) {
	return nil, nil
}
