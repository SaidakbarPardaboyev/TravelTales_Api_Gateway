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
			log.Printf("authorization header is required")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "http.StatusBadRequest",
				"massege": "authorization header is required",
			})
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			log.Printf("error while validating token: %s", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "http.StatusBadRequest",
				"massege": fmt.Sprintf("error while validating token: %s", err),
			})
			return
		}

		// claims, err := token.ExtractClaims(auth)
		// if err != nil || !valid {
		// 	ctx.AbortWithStatusJSON(http.StatusBadRequest,
		// 		fmt.Errorf("invalid token claims: %s", err))
		// 	return
		// }

		// ctx.Set("claims", claims)
		ctx.Next()
	}
}
