package main

import (
	"log"
	"os"

	"github.com/first/routes"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// 初始化 .env 的配置，将 .env 中的配置加载到 Go 的 env 环境中
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	port := os.Getenv("LISTEN_PORT")

	// 路由
	routes.RouteInit(*app)

	app.Run(iris.Addr(port))
}
