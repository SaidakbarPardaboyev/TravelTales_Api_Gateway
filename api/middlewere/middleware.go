package middlewere

import (
	"fmt"
	"log"
	"net/http"
	"travel/api/token"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			log.Printf("id not found from url params")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "http.StatusBadRequest",
				"massege": "id not found from url params",
			})
			return
		}

		valid, err := token.ValidateToken(auth)
		if !valid || err != nil {
			log.Printf("error while validating token: %s", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "http.StatusBadRequest",
				"massege": fmt.Sprintf("error while validating token: %s", err),
			})
			return
		}

		ctx.Next()
	}
}
