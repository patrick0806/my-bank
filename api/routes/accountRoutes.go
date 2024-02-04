package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAccountsRoutes(rg *gin.RouterGroup) {
	accountRouter := rg.Group("/accounts")

	accountRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "{\"amigo\":\"estou aqui\"}")
	})
}
