package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/dtos"
)

type AccountController struct {
	CreateAccountUseCase    usecases.CreateAccountUseCase
	FindAccountByCPFUseCase usecases.FindAccountByCPFUseCase
}

func NewControllerAccount(
	createAccountUseCase usecases.CreateAccountUseCase,
	findAccountByCPFUseCase usecases.FindAccountByCPFUseCase) *AccountController {
	return &AccountController{
		CreateAccountUseCase:    createAccountUseCase,
		FindAccountByCPFUseCase: findAccountByCPFUseCase,
	}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description My bank new client and account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body usecases.CreateAccountRequestDTO true "Create account parameters"
// @Success 201 {object} usecases.CreateAccountResponseDTO
// @Failure      400  {object}  dtos.HttpErrorDTO
// @Failure      409  {object}  dtos.HttpErrorDTO
// @Failure      500  {object}  dtos.HttpErrorDTO
// @Router       /accounts [post]
func (ac *AccountController) CreateAccount(ctx *gin.Context) {
	requestDTO := usecases.CreateAccountRequestDTO{}
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}

	createdAccountDTO, err := ac.CreateAccountUseCase.Execute(requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, createdAccountDTO)
}

// FindAccountByCPF godoc
// @Summary Find account by CPF
// @Description Find account using the Brazilian individual taxpayer registry identification number (CPF)
// @Tags Accounts
// @Accept json
// @Produce json
// @Param cpf path string true "Brazilian individual taxpayer registry identification number (CPF)"
// @Success 200 {object} usecases.FindAccountByCPFResponseDTO "Account found successfully"
// @Failure 400 {object} dtos.HttpErrorDTO
// @Failure 404 {object} dtos.HttpErrorDTO
// @Failure 500 {object} dtos.HttpErrorDTO
// @Router /accounts/{cpf} [get]
func (ac *AccountController) FindAccountByCPF(ctx *gin.Context) {
	createdAccountDTO, err := ac.FindAccountByCPFUseCase.Execute(ctx.Param("cpf"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}
	ctx.JSON(http.StatusOK, createdAccountDTO)
}
