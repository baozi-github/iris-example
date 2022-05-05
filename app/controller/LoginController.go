package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	"fmt"
)

// 登录控制器
type LoginController struct {
	BaseController Controller
}

type LoginData struct {
	Name     string `json:"name" validate:"required" comment:"名称"`
	Password string `json:"password" validate:"required" comment:"密码"`
}

type ErrorMessage struct {
}

// var ErrorMessage map[string]string

func (c *LoginController) GetLogin(ctx iris.Context) {
	var loginData LoginData
	err := ctx.ReadJSON(&loginData)
	err = validator.New().Struct(loginData)
	ctx.JSON(loginData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field())

			break
		}

		ctx.JSON("error")
		ctx.JSON(err)

		return
	}
	ctx.JSON("error111111111")
	return
	// return common.ReturnSuccess(200, "操作成功", map[string]string{"name": "hahahah"})
}
