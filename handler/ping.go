package handler

import "github.com/gin-gonic/gin"

// PingJSON は疎通確認用
func PingJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "good",
	})
	return
}
