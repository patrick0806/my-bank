package queues

import (
	"fmt"

	"github.com/patrick0806/my-bank/internal/entities"
)

var signalChan = make(chan struct{})

type TransactionQueue struct {
	transactions []entities.Transaction
}

func NewTransactionQueue() *TransactionQueue {
	return &TransactionQueue{
		transactions: make([]entities.Transaction, 0),
	}
}

func (q *TransactionQueue) Enqueue(t entities.Transaction) {
	q.transactions = append(q.transactions, t)
	signalChan <- struct{}{}
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
			<-signalChan
			continue
		}
		// Processar a transação aqui...
		fmt.Printf("Processing transaction ID: %s\n", transaction.Id.String())
	}
}
