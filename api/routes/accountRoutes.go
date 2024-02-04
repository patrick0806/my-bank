package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/controllers"
	"github.com/patrick0806/my-bank/internal/usecases"
)

func AddAccountsRoutes(rg *gin.RouterGroup) {
	accountRouter := rg.Group("/accounts")
	accountController := controllers.NewControllerAccount(usecases.CreateAccountUseCase{})

	accountRouter.POST("/", accountController.CreateAccount)
}
