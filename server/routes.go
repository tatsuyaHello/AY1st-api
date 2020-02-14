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

	public := base.Group("/", CacheMiddleware())

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
		withUser.GET("me", handler.GetMe)
		withUser.GET("/users/:user-id", RequirePathParam("user-id"), handler.GetUser)
		public.POST("/users", handler.PostUser)
		withUser.PUT("/users/:user-id", RequirePathParam("user-id"), handler.PutUser)
		withUser.DELETE("/users/:user-id", RequirePathParam("user-id"), handler.DeleteUser)
	}

	// Book
	{
		base.POST("/books", handler.PostBooks)
	}

}
