package main

import (
	"AY1st/server"
	"AY1st/util"
	"os"
)

func main() {

	util.LoadEnv()
	util.InitLogger()
	err := server.Start()
	if err != nil {
		util.GetLogger().Errorln(err)
		os.Exit(1)
	}

	// router := gin.Default()

	// // ルーティングの登録を以下で行う
	// router.GET("ping/json", handler.PingJSON)

	// // err := godotenv.Load()
	// // if err != nil {
	// // 	log.Fatal("error loading .env file")
	// // }

	// // log.Fatal(router.Run(":" + os.Getenv("SERVER_PORT")))
	// log.Fatal(router.Run(":" + "8080"))
}
