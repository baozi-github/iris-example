package routes

import (
	"github.com/first/app/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RouteInit(app iris.Application) {
	mvc.New(app).Handle(new(controller.ExampleController))
}
