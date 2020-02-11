package server

import (
	"AY1st/handler"
	"AY1st/registry"
	"AY1st/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserMiddleware 認証したユーザー情報を取得する
func UserMiddleware() gin.HandlerFunc {
	logger := util.GetLogger()
	return func(c *gin.Context) {
		err := UserHandler(c)
		if err != nil {
			er := handler.NewErrorResponse("401", handler.ErrorAuth, err.Error())
			logger.Debug(er)
			c.AbortWithStatusJSON(http.StatusUnauthorized, er)
			return
		}

		c.Next()
	}
}

// OptionalUserMiddleware 認証していればユーザー情報を取得する
func OptionalUserMiddleware() gin.HandlerFunc {
	logger := util.GetLogger()
	return func(c *gin.Context) {
		_, ok := c.Get("email")
		if ok {
			err := UserHandler(c)
			if err != nil {
				er := handler.NewErrorResponse("401", handler.ErrorAuth, err.Error())
				logger.Debug(er)
				c.AbortWithStatusJSON(http.StatusUnauthorized, er)
				return
			}
		}

		c.Next()
	}
}

// UserHandler は認証情報からユーザー取得を行う
func UserHandler(c *gin.Context) error {

	servicer := c.MustGet(registry.ServiceKey).(registry.Servicer)
	usersService := servicer.NewUsers()

	// ID Tokenのemailはエンドユーザー身元識別子
	email := c.MustGet("email").(string)
	user, err := usersService.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("cannot find your identity code")
	}

	sub := c.MustGet("sub").(string)
	_, err = usersService.GetUserID(sub)

	if err != nil {
		err := usersService.AddIdentity(user.ID, sub)
		if err != nil {
			return err
		}
	}

	c.Set("user", user)
	return nil
}
