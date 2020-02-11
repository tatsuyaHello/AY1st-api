package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExposeHeader は追加ヘッダに公開属性を付与します
func ExposeHeader(c *gin.Context, headerKey string) {
	exposeList := c.Writer.Header().Get("Access-Control-Expose-Headers")
	var updated string
	if exposeList == "" {
		updated = headerKey
	} else {
		updated = exposeList + "," + headerKey
	}

	c.Header("Access-Control-Expose-Headers", updated)
}

// SetHeaderXTotalCount は X-Total-Count ヘッダを追加します
func SetHeaderXTotalCount(c *gin.Context, count int) {
	ExposeHeader(c, "x-total-count")
	c.Header("x-total-count", strconv.Itoa(count))
}
