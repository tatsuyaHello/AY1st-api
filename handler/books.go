package handler

import (
	"AY1st/model"
	"AY1st/registry"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostBooks は投稿の新規登録
func PostBooks(c *gin.Context) {
	// user := c.MustGet("user").(*model.User)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	booksService := servicer.NewBooks()

	input := model.BookActionInput{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, err.Error()))
		return
	}

	created, err := booksService.Create(user.ID, &input)
	// created, err := booksService.Create(1, &input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, created)
}
