package server

import "github.com/gin-gonic/gin"

// build command
// go build -ldflags "-X package/path/to/server.version=$(git rev-parse --short HEAD)"

// version Git Commit Hash
var version = "unknown"

// VersionMiddleware は Version 埋め込みます
func VersionMiddleware(showVersion bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if showVersion {
			c.Writer.Header().Set("x-api-version", version)
		}
	}
}
