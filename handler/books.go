package handler

import (
	"AY1st/model"
	"AY1st/registry"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetBook は本の取得
func GetBook(c *gin.Context) {
	bookID := c.MustGet("book-id").(uint64)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	booksService := servicer.NewBooks()

	book, err := booksService.GetOne(bookID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, book)
}

// PostBook は本の新規登録
func PostBook(c *gin.Context) {

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	booksService := servicer.NewBooks()

	input := model.BookBody{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, err.Error()))
		return
	}

	created, err := booksService.Create(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, created)
}
