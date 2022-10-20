package api

import (
	request "github.com/HolisChen/tiny-blog/model/request/user"
	"github.com/HolisChen/tiny-blog/model/response"
	"github.com/HolisChen/tiny-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
登录方法
*/

// 登录方法 godoc
// @Summary 登录
// @Schemes
// @Description 执行登录返回token
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body request.UserLoginRequest true "登录请求"
// @Success 200 {object} response.LoginResponse "Response"
// @Router /user/login [post]
func Login(c *gin.Context) {
	loginReq := &request.UserLoginRequest{}
	err := c.ShouldBind(loginReq)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err))
		return
	}
	res, err := service.DoLogin(loginReq)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

/**
注册方法
*/
// 注册 godoc
// @Summary 注册
// @Schemes
// @Description 用户注册并执行自动登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body request.UserRegisterRequest true "注册请求"
// @Success 200 {object} response.LoginResponse "Response"
// @Router /user/register [post]
func Register(c *gin.Context) {
	req := &request.UserRegisterRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err))
		return
	}
	err = service.DoRegister(req)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err))
		return
	}
	c.JSON(http.StatusOK, response.OK())

}

/**
登出方法
*/
func Logout(c *gin.Context) {

}

/**
重置密码
*/
func ResetPwd(c *gin.Context) {

}
