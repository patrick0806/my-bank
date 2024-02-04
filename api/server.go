package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/api/routes"
)

func Run(port string) {
	router := gin.Default()
	getRoutes(router)
	router.Run(port)
}

func getRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	routes.AddAccountsRoutes(v1)
}
