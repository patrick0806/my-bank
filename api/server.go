package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/routes"
	"github.com/patrick0806/my-bank/pkg/queues"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(port string, db *sql.DB, queue *queues.TransactionQueue) {
	router := gin.Default()
	// add swagger
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	getRoutes(router, db, queue)
	router.Run(port)
}

func getRoutes(router *gin.Engine, db *sql.DB, queue *queues.TransactionQueue) {
	v1 := router.Group("api/v1")
	routes.AddAuthRoutes(v1, db)
	routes.AddAccountsRoutes(v1, db)
	routes.AddTransacationRoutes(v1, db, queue)
}
