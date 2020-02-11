package server

import (
	"fmt"
	"net/http"
	"strconv"

	"AY1st/model"

	"github.com/gin-gonic/gin"
)

// RequirePathParam parses PathParameter as uint64 by given param and then sets it to gin Context.
func RequirePathParam(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param(param), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				model.NewErrorResponseWithCode(model.ErrorResourceNotFound, fmt.Sprintf("%v must be a positive number", param)),
			)
			return
		}
		c.Set(param, id)
		c.Next()
	}
}
