package routes

import "github.com/kataras/iris/v12"

func RouteInit(app iris.Application) {
	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
}
