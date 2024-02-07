package queues

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/pkg/datasources"
)

type TransactionQueue struct {
	transactions      []entities.Transaction
	accountDataSource datasources.AccountDatasource
}

func NewTransactionQueue(db *sql.DB) *TransactionQueue {
	accountDatasource := datasources.AccountDatasource{
		DB: db,
	}
	return &TransactionQueue{
		transactions:      make([]entities.Transaction, 0),
		accountDataSource: accountDatasource,
	}
}

func (q *TransactionQueue) Enqueue(t entities.Transaction) {
	q.transactions = append(q.transactions, t)
}

func (q *TransactionQueue) Dequeue() (entities.Transaction, bool) {
	if len(q.transactions) == 0 {
		return entities.Transaction{}, false
	}

	transaction := q.transactions[0]
	q.transactions = q.transactions[1:]
	return transaction, true
}

func (q *TransactionQueue) ProcessTransactions() {
	for {
		transaction, ok := q.Dequeue()
		if !ok {
			//Olhar como fazer isso funcionar com channels e rang se posivel
			// Se a fila estiver vazia, esperar por uma nova transação
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Printf("Processing transaction ID: %s\n", transaction.Id.String())

		accountOrigin, err := q.accountDataSource.FindById(transaction.AccountOrigin)
		if err != nil || err == nil && accountOrigin == nil {
			continue
		}

		if transaction.Value < accountOrigin.Balance {
			println("Fail to continue you need more balance")
			continue
		}

		accountDestiny, err := q.accountDataSource.FindById(transaction.AccountDestiny)
		if err != nil || err == nil && accountDestiny == nil {
			continue
		}

		accountOrigin.Balance = accountOrigin.Balance - transaction.Value
		accountDestiny.Balance = accountDestiny.Balance + transaction.Value

		q.accountDataSource.UpdateBalance(accountDestiny.Id.String(), accountDestiny.Balance)
		q.accountDataSource.UpdateBalance(accountOrigin.Id.String(), accountOrigin.Balance)
	}
}
