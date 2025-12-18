package middleware

import (
	"fmt"
	"log"
	"myapp/pkg/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Panic] %v\n%s", err, string(debug.Stack()))
				response.HandleError(c, fmt.Errorf("system panic: %v", err))
				c.Abort()
			}
		}()
		c.Next()
	}
}
