package handler

import (
	"AY1st/model"
	"AY1st/registry"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostPost は投稿の新規登録
func PostPost(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	postsService := servicer.NewPosts()

	input := model.PostInput{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, err.Error()))
		return
	}

	created, err := postsService.Create(user.ID, &input)
	// created, err := postsService.Create(1, &input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, created)
}

// GetPost は投稿の新規登録
func GetPost(c *gin.Context) {
	postID := c.MustGet("post-id").(uint64)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	postsService := servicer.NewPosts()

	post, err := postsService.GetOne(postID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, post)
}

// GetPostAll は投稿の新規登録
func GetPostAll(c *gin.Context) {

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	postsService := servicer.NewPosts()

	posts, err := postsService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, posts)
}

// DeletePost は投稿の削除
func DeletePost(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	// postIDはユーザと本の関連性に関するID
	postID := c.MustGet("post-id").(uint64)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	postsService := servicer.NewPosts()

	post, err := postsService.GetOne(postID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorUnknown, err.Error()))
		return
	}

	// ユーザーの整合性を確認
	if post.UserID != user.ID {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, "user-id is not your own id"))
		return
	}

	err = postsService.Delete(postID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorLimitExceeded, err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}
