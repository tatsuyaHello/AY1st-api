package main

import (
	"log"

	"AY1st/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// ルーティングの登録を以下で行う
	router.GET("ping/json", handler.PingJSON)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error loading .env file")
	// }

	// log.Fatal(router.Run(":" + os.Getenv("SERVER_PORT")))
	log.Fatal(router.Run(":" + "8080"))
}
