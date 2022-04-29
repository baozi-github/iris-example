package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	// 初始化 .env 的配置，将 .env 中的配置加载到 Go 的 env 环境中
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatalln(err)
	// 	os.Exit(1)
	// }

	// port := os.Getenv("LISTEN_PORT")

	config := iris.WithConfiguration(iris.YAML("./iris.yml"))
	println(config)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8082"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
