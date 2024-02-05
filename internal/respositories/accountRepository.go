package respositories

import "github.com/patrick0806/my-bank/internal/entities"

type AccountRepository interface {
	FindByCPF(cpf string) (*entities.Account, error)
	CreateAccount(account *entities.Account) error
}
