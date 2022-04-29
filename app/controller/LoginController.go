package controller

import (
	"github.com/kataras/iris/v12"
)

// 登录控制器
type LoginController struct {
	BaseController Controller
}

type LoginData struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *LoginController) GetLogin(ctx iris.Context) {
	var loginData LoginData
	err := ctx.ReadJSON(&loginData)
	if err != nil {
		ctx.JSON("error")
		return
	}
	ctx.JSON("error111111111")
	return
	// return common.ReturnSuccess(200, "操作成功", map[string]string{"name": "hahahah"})
}
