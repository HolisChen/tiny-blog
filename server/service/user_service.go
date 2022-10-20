package service

import (
	"errors"
	"github.com/HolisChen/tiny-blog/model"
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
	encryptPassword := tools.GenerateMd5(req.Password + user.MailAddress + user.LoginId)
	if user.Password != encryptPassword {
		return nil, errors.New("用户密码错误~")
	}
	//check status
	if user.Status != 0 {
		return nil, errors.New("用户已被锁定，禁止登录~")
	}
	//登录成功更新最后登录时间
	go updateLastLoginTime(user.ID)

	return &response.LoginResponse{
		JWT: tools.GenerateJwt(user),
	}, nil
}

/**
注册用户
*/
func DoRegister(req *request.UserRegisterRequest) error {
	if req.Password != req.RePassword {
		return errors.New("两次密码不一致~")
	}

	user := repository.FindUserByLoginId(req.LoginId)
	if user != nil {
		return errors.New("当前登录ID已被注册~")
	}
	user = repository.FindByMail(req.Mail)
	if user != nil {
		return errors.New("当前邮箱已注册~")
	}
	//执行插入
	user = &model.Author{
		Name:          req.Name,
		LoginId:       req.LoginId,
		Password:      tools.GenerateMd5(req.Password + req.Mail + req.LoginId),
		LastLoginTime: nil,
		Level:         "初入江湖",
		Status:        0,
		MailAddress:   req.Mail,
	}
	err := repository.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func updateLastLoginTime(userId uint) {
	repository.UpdateLastLoginTime(userId)
}
