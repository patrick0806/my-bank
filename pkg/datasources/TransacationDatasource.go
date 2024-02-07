package datasources

import (
	"database/sql"

	"github.com/patrick0806/my-bank/internal/entities"
)

type TransactionDatasource struct {
	DB *sql.DB
}

func NewTransactionDatasource(db *sql.DB) *TransactionDatasource {
	return &TransactionDatasource{
		DB: db,
	}
}

func (td TransactionDatasource) CreateNewTransacation(entities.Transaction) error {
	return nil
}
