package queues

import (
	"sync"

	"github.com/patrick0806/my-bank/internal/entities"
)

type TransactionQueue struct {
	transactions []entities.Transaction
	mu           sync.Mutex
}

func NewTransactionQueue() *TransactionQueue {
	return &TransactionQueue{
		transactions: make([]entities.Transaction, 0),
	}
}

func (q *TransactionQueue) Enqueue(t entities.Transaction) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.transactions = append(q.transactions, t)
}

func (q *TransactionQueue) Dequeue() (entities.Transaction, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.transactions) == 0 {
		return entities.Transaction{}, false
	}

	transaction := q.transactions[0]
	q.transactions = q.transactions[1:]
	return transaction, true
}
