package main

import (
	"io"
	"log"
	"os"

	"AY1st/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// logファイルの設定
	logfile, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("cannot open logfile:%v", err)
	}
	defer logfile.Close()
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.Print("ok?")

	router := gin.Default()

	// ルーティングの登録を以下で行う
	router.GET("ping/json", handler.PingJSON)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	log.Fatal(router.Run(":" + os.Getenv("SERVER_PORT")))
}
