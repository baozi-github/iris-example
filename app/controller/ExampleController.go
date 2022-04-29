package controller

type ExampleController struct {
}

func (c *ExampleController) Get() string {
	return "请求成功"
}
