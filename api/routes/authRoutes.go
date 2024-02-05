package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/controllers"
	"github.com/patrick0806/my-bank/internal/usecases"
	"github.com/patrick0806/my-bank/pkg/datasources"
)

func AddAuthRoutes(rg *gin.RouterGroup, db *sql.DB) {
	authRouter := rg.Group("/auth")
	accountDatasource := datasources.AccountDatasource{
		DB: db,
	}
	signInUseCase := usecases.SignInUseCase{AccountRepository: &accountDatasource}
	authController := controllers.NewAuthController(signInUseCase)

	authRouter.POST("/", authController.SignIn)
}
