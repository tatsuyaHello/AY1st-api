package server

import (
	"github.com/gin-gonic/gin"
)

// CacheMiddleware cache middleware
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("cache-control", "no-cache")
		c.Next()
	}
}
