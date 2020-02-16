package handler

import (
	"AY1st/registry"
	"net/http"

	"github.com/gin-gonic/gin"
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
