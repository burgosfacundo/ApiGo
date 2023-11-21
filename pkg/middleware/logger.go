package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verb := ctx.Request.Method
		time := time.Now()
		url := ctx.Request.URL
		var size int

		ctx.Next()

		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		fmt.Printf("\nPath:%s\n	Verb: %s\n	Time: %v\n	Size: %d\n", url, verb, time, size)
	}
}
