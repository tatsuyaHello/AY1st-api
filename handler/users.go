package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"AY1st/model"
	"AY1st/registry"

	"github.com/gin-gonic/gin"
)

// GetMe is
func GetMe(c *gin.Context) {
	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	// subからユーザを取得
	sub, ok := c.Get("sub")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	me, err := usersService.GetMe(sub.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, me)
}

// GetUser は指定ユーザを取得
func GetUser(c *gin.Context) {
	userID := c.MustGet("user-id").(uint64)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	user, err := usersService.GetOne(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, NewErrorResponse("404", ErrorNotFound, fmt.Sprintf("no result for user-id = %v", userID)))
		return
	}
	c.JSON(http.StatusOK, user)
}

// PostUser は新規ユーザー登録
func PostUser(c *gin.Context) {
	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	input := model.UserCreateInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponseDetailed("400", ErrorParam, "request body mismatch", err.Error()))
		return
	}

	created, err := usersService.Create(&input)
	if err != nil {
		errCode, _ := err.(*model.WrappedError)
		if IsUserSubDupulicateError(err) {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(string(errCode.Code), ErrorParam, errCode.Message))
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, err.Error()))
		return
	}

	c.JSON(http.StatusOK, created)
}

// PutUser はユーザー更新
func PutUser(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	userID := c.MustGet("user-id").(uint64)

	// ユーザーの整合性を確認
	if user.ID != userID {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, "user-id is not your own id"))
		return
	}

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	input := &model.UserUpdateInput{}
	err := c.ShouldBindWith(input, binding.JSON)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponseDetailed("400", ErrorParam, "request body mismatch", err.Error()))
		return
	}

	updated, err := usersService.Update(userID, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, err.Error()))
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteUser はユーザー情報削除
func DeleteUser(c *gin.Context) {

	user := c.MustGet("user").(*model.User)
	userID := c.MustGet("user-id").(uint64)

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	// ユーザーの整合性を確認
	if user.ID != userID {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorParam, "user-id is not your own id"))
		return
	}

	err := usersService.Delete(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("400", ErrorLimitExceeded, err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}
