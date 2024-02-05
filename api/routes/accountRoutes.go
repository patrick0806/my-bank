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
	accountRouter := rg.Group("/accounts")
	accountController := controllers.NewControllerAccount(usecases.CreateAccountUseCase{
		AccountRepository: &accountDatasource,
	})

	accountRouter.POST("/", accountController.CreateAccount)
}
