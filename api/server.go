package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(port string) {
	router := gin.Default()
	// add swagger
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	getRoutes(router)
	router.Run(port)
}

func getRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	routes.AddAccountsRoutes(v1)
}
