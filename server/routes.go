package server

import (
	"net/http"

	"AY1st/handler"
	"AY1st/service"

	"github.com/gin-gonic/gin"
)

func defineRoutes(r gin.IRouter, authenticator Authenticator, userService service.UsersInterface, cacheMaxAge int64) {

	// 非ログイン状態で使用したい場合はこちらを使用する
	public := r.Group("/", CacheMiddleware())

	// ログイン状態で使用したい場合はこちらを使用する
	withUser := r.Group("/",
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
		r.GET("ping/json", handler.PingJSON)
		r.GET("health", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		r.GET("db/health", handler.GetHealth)
	}

	// User
	{
		public.POST("/signup", handler.PostUser)
		withUser.GET("/me", handler.GetMe)
		withUser.GET("/users/:user-id", RequirePathParam("user-id"), handler.GetUser)
		withUser.PATCH("/users/:user-id", RequirePathParam("user-id"), handler.PutUser)
		withUser.DELETE("/users/:user-id", RequirePathParam("user-id"), handler.DeleteUser)

	}

	// Post
	{
		withUser.POST("/posts", handler.PostPost)
		withUser.GET("/posts", handler.GetPostAll)
		withUser.GET("/posts/:post-id", RequirePathParam("post-id"), handler.GetPost)
		withUser.PATCH("/posts/:post-id", RequirePathParam("post-id"), handler.PutPost)
		withUser.DELETE("/posts/:post-id", RequirePathParam("post-id"), handler.DeletePost)
		withUser.GET("/posts_users/:user-id", RequirePathParam("user-id"), handler.GetPostOfUser)
	}

	// Book
	{
		withUser.GET("/books/:book-id", RequirePathParam("book-id"), handler.GetBook)
	}

}
