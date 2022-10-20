package response

import "net/http"

type BaseResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) BaseResponse {
	return BaseResponse{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: data,
	}
}

func OK() BaseResponse {
	return Success(nil)
}

func Fail(msg string) BaseResponse {
	return BaseResponse{
		Code: -1,
		Msg:  msg,
	}
}

func Error(err error) BaseResponse {
	return BaseResponse{
		Code: -1,
		Msg:  err.Error(),
	}
}
