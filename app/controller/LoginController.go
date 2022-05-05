package controller

import (
	"github.com/first/common"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// 登录控制器
type LoginController struct {
	BaseController Controller
}

// 接收数据以及校验数据
type LoginData struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// 定义错误信息
var LoginDataErrorMessage = map[string]string{
	"Name_required":     "名称是必填的",
	"Password_required": "密码是必填的",
}

// 登录接口
func (c *LoginController) GetLogin(ctx iris.Context) common.ReturnData {
	var loginData LoginData
	err := ctx.ReadJSON(&loginData)
	err = validator.New().Struct(loginData)
	ctx.JSON(loginData)
	if err != nil {
		errMsg := ""
		for _, err := range err.(validator.ValidationErrors) {
			errMsgKey := err.Field() + "_" + err.Tag()
			errMsg = LoginDataErrorMessage[errMsgKey]
			break
		}
		ctx.JSON(errMsg)

		return common.ReturnFail(200, errMsg)
	}
	return common.ReturnSuccess(200, "请求成功", "")
}
