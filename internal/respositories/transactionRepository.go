package respositories

import "github.com/patrick0806/my-bank/internal/entities"

type TransacationRepository interface {
	CreateNewTransacation(entities.Transaction) error
}
