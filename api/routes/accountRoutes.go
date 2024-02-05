package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/controllers"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/datasources"
)

func AddAccountsRoutes(rg *gin.RouterGroup, db *sql.DB) {
	accountDatasource := datasources.AccountDatasource{DB: db}
	createAccountUseCase := usecases.CreateAccountUseCase{AccountRepository: &accountDatasource}
	findAccountByCPF := usecases.FindAccountByCPFUseCase{AccountRepository: &accountDatasource}

	accountRouter := rg.Group("/accounts")
	accountController := controllers.NewControllerAccount(createAccountUseCase, findAccountByCPF)

	accountRouter.POST("/", accountController.CreateAccount)
	accountRouter.GET("/:cpf", accountController.FindAccountByCPF)
}
