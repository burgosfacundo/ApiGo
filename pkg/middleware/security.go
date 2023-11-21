package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Auth is a function that authenticates token requests
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("token")
		tokenEnv := os.Getenv("TOKEN_ENV")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		} else {
			ctx.Next()
		}
	}
}
