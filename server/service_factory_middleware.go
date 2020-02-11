package server

import (
	"AY1st/registry"

	"github.com/gin-gonic/gin"
)

// ServiceKeyMiddleware provides the service registory
func ServiceKeyMiddleware(si registry.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(registry.ServiceKey, si)
		c.Next()
	}
}
