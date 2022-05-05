# iris-example
# iris框架代码示例


### 安装过程
```go
git clone git@github.com:baozi-github/iris-example.git
cd iris-example
go build
```


### 数据统一返回
```go
type ReturnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 操作正确返回
func ReturnSuccess(code int, message string, data interface{}) ReturnData {
	return ReturnData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// 操作失败返回 data为空
func ReturnFail(code int, message string) ReturnData {
	return ReturnData{
		Code:    code,
		Message: message,
		Data:    "",
	}
}
```


### 数据校验功能
```go
go get github.com/go-playground/validator/v10
```

### 自定义错误信息
```go
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
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// 获取map中的key 如果key不存在 统一定义为参数错误
			errMsgKey := err.Field() + "_" + err.Tag()
			errMsg, ok := LoginDataErrorMessage[errMsgKey]
			if !ok {
				errMsg = "参数错误"
			}
			// 只返回第一条错误
			return common.ReturnFail(200, errMsg)
		}
		return common.ReturnFail(200, "参数错误")
	}
	return common.ReturnSuccess(200, "请求成功", "")
}
```



