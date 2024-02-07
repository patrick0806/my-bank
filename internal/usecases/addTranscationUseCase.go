package usecases

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/internal/respositories"
	"github.com/patrick0806/my-bank/pkg/entity"
)

type AddTransacationUseCase struct {
	TransactionRepository respositories.TransacationRepository
}

type TransactionRequestDTO struct {
	Value          float64 `json:"value"`
	AccountOrigin  string  `json:"accountOrigin"`
	AccountDestiny string  `json:"accountDestiny"`
}

var (
	transactionQueue = make(chan entities.Transaction, 100) // Buffer size of  100
	wg               sync.WaitGroup
)

func NewAddTransacationUseCase(repository respositories.TransacationRepository) *AddTransacationUseCase {
	return &AddTransacationUseCase{
		TransactionRepository: repository,
	}
}

func (us *AddTransacationUseCase) Execute(transcationDTO TransactionRequestDTO) error {
	transcation := entities.Transaction{
		Id:             entity.NewID(),
		Value:          transcationDTO.Value,
		AccountOrigin:  transcationDTO.AccountOrigin,
		AccountDestiny: transcationDTO.AccountDestiny,
	}
	fmt.Printf("Add transaction in queue: %v", transcation)
	wg.Add(1)
	transactionQueue <- transcation
	go processTransactions()
	return nil
}

func processTransactions() {
	defer fmt.Printf("Saindo do process")
	defer wg.Done()
	for tx := range transactionQueue {
		// Processar a transação aqui...
		time.Sleep(time.Second * 2)
		println("Processing transaction ID:", tx.Id.String())
		// ...
	}
}
