package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/dtos"
)

type TransactionController struct {
	AddTransacationUseCase usecases.AddTransacationUseCase
}

func NewTransactionController(addTransaction usecases.AddTransacationUseCase) *TransactionController {
	return &TransactionController{
		AddTransacationUseCase: addTransaction,
	}
}

// Add Transaction godoc
// @Summary Create Transacation
// @Description Add a bank transacation
// @Tags Transactions
// @Accept json
// @Produce json
// @Param account body usecases.TransactionRequestDTO true "Transaction details"
// @Success 200
// @Failure      400  {object}  dtos.HttpErrorDTO
// @Failure      500  {object}  dtos.HttpErrorDTO
// @Router       /transactions [post]
// @Security ApiKeyAuth
func (ac *TransactionController) AddTransacation(ctx *gin.Context) {
	requestDTO := usecases.TransactionRequestDTO{}
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}

	err = ac.AddTransacationUseCase.Execute(requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}

	ctx.Status(http.StatusOK)
}
