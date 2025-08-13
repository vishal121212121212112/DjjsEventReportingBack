package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Request-Id") == "" {
			c.Writer.Header().Set("X-Request-Id", uuid.NewString())
		}
		c.Next()
	}
}
