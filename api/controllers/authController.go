package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/dtos"
)

type AuthController struct {
	SignInUseCase usecases.SignInUseCase
}

func NewAuthController(signInUseCase usecases.SignInUseCase) *AuthController {
	return &AuthController{
		SignInUseCase: signInUseCase,
	}
}

// Login godoc
// @Summary Login
// @Description Login in api with cpf and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param account body usecases.SignRequestDTO true "Login params"
// @Success 201 {object} usecases.SignResponseDTO
// @Failure      400  {object}  dtos.HttpErrorDTO
// @Failure      500  {object}  dtos.HttpErrorDTO
// @Router       /auth [post]
func (ac *AuthController) SignIn(ctx *gin.Context) {
	requestDTO := usecases.SignRequestDTO{}
	err := ctx.ShouldBindJSON(&requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}

	accessToken, err := ac.SignInUseCase.Execute(requestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.HttpErrorDTO{
			StatusCode: 400,
			Reason:     err.Error(),
			Timestamp:  time.Now(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, accessToken)
}
