package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/controllers"
	"github.com/patrick0806/my-bank/api/middlewares"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/datasources"
)

func AddTransacationRoutes(rg *gin.RouterGroup, db *sql.DB) {
	transactionRouter := rg.Group("/transactions")
	transactionDataSource := datasources.TransactionDatasource{
		DB: db,
	}
	addTransactionUseCase := usecases.AddTransacationUseCase{
		TransactionRepository: transactionDataSource,
	}
	transactionController := controllers.NewTransactionController(addTransactionUseCase)

	transactionRouter.POST("/", middlewares.ValidateAccessToken, transactionController.AddTransacation)
}
