package usecases

import (
	"fmt"

	"github.com/patrick0806/my-bank/internal/entities"
	"github.com/patrick0806/my-bank/internal/respositories"
	"github.com/patrick0806/my-bank/pkg/entity"
	"github.com/patrick0806/my-bank/pkg/queues"
)

type AddTransacationUseCase struct {
	TransactionRepository respositories.TransacationRepository
	Queue                 *queues.TransactionQueue
}

type TransactionRequestDTO struct {
	Value          float64 `json:"value"`
	AccountOrigin  string  `json:"accountOrigin"`
	AccountDestiny string  `json:"accountDestiny"`
}

func NewAddTransacationUseCase(repository respositories.TransacationRepository, queue *queues.TransactionQueue) *AddTransacationUseCase {
	return &AddTransacationUseCase{
		TransactionRepository: repository,
		Queue:                 queue,
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
	us.Queue.Enqueue(transcation)
	return nil
}
