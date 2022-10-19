package service

import (
	"errors"
	request "github.com/HolisChen/tiny-blog/model/request/user"
	response "github.com/HolisChen/tiny-blog/model/response/login"
	"github.com/HolisChen/tiny-blog/repository"
	"github.com/HolisChen/tiny-blog/tools"
)

func DoLogin(req *request.UserLoginRequest) (*response.LoginResponse, error) {
	user := repository.FindUserByLoginId(req.LoginId)
	if user == nil {
		return nil, errors.New("用户不存在~")
	}
	//check password
	if user.Password != req.Password {
		return nil, errors.New("用户密码错误~")
	}
	//check status
	if user.Status != 0 {
		return nil, errors.New("用户已被锁定，禁止登录~")
	}
	return &response.LoginResponse{
		JWT: tools.GenerateJwt(user),
	}, nil
}
