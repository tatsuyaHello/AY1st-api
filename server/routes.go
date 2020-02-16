package server

import (
	"net/http"
	"os"

	"AY1st/handler"
	"AY1st/service"

	"github.com/gin-gonic/gin"
)

func defineRoutes(r gin.IRouter, authenticator Authenticator, userService service.UsersInterface, cacheMaxAge int64) {

	base := r.Group("/" + os.Getenv("ENVCODE"))

	// 非ログイン状態で使用したい場合はこちらを使用する
	public := base.Group("/", CacheMiddleware())

	// ログイン状態で使用したい場合はこちらを使用する
	withUser := base.Group("/",
		AuthMiddleware(authenticator),
		UserMiddleware(),
		CacheMiddleware(),
	)

	// withOptionalUser := base.Group("/",
	// 	OptionalAuthMiddleware(authenticator),
	// 	OptionalUserMiddleware(),
	// 	CacheMiddleware(),
	// )

	// Health Check
	{
		base.GET("ping/json", handler.PingJSON)
		base.GET("health", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		base.GET("db/health", handler.GetHealth)
	}

	// User
	{
		public.POST("/signup", handler.PostUser)
		withUser.GET("me", handler.GetMe)
		withUser.GET("/users/:user-id", RequirePathParam("user-id"), handler.GetUser)
		withUser.PATCH("/users/:user-id", RequirePathParam("user-id"), handler.PutUser)
		withUser.DELETE("/users/:user-id", RequirePathParam("user-id"), handler.DeleteUser)

	}

	// Post
	{
		withUser.POST("/posts", handler.PostPost)
		withUser.GET("/posts", handler.GetPostAll)
		withUser.GET("/posts/:post-id", RequirePathParam("post-id"), handler.GetPost)
		withUser.PUT("/posts/:post-id", RequirePathParam("post-id"), handler.PutPost)
		withUser.DELETE("/posts/:post-id", RequirePathParam("post-id"), handler.DeletePost)
	}

}
