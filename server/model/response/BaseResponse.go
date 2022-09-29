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
