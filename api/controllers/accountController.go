package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/dtos"
)

type AccountController struct {
	CreateAccountUseCase usecases.CreateAccountUseCase
}

func NewControllerAccount(createAccountUseCase usecases.CreateAccountUseCase) *AccountController {
	return &AccountController{
		CreateAccountUseCase: createAccountUseCase,
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
