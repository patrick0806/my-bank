package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrick0806/my-bank/pkg/dtos"
	"github.com/patrick0806/my-bank/pkg/utils"
)

func ValidateAccessToken(ctx *gin.Context) {

	accessToken := ctx.GetHeader("accessToken")
	err := utils.VerifyToken(accessToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, &dtos.HttpErrorDTO{
			StatusCode: http.StatusUnauthorized,
			Reason:     "Invalid access token",
			Timestamp:  time.Now(),
		})
		return
	}

	ctx.Next()

}
