package middlewere

import (
	"context"

	"github.com/gin-gonic/gin"
)

func () Middleware(ctx *gin.Context) *gin.HandlerFunc {
	return &gin.HandlersChain{

		auth := ctx.GetHeader("Authorization")

		if auth == "" {

		}
	}
}