package common

type ReturnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnSuccess(code int, message string, data interface{}) ReturnData {
	return ReturnData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
