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
	}
	res, err := service.DoLogin(loginReq)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err))
	}
	c.JSON(http.StatusOK, response.Success(res))
}

/**
注册方法
*/
func Register(c *gin.Context) {

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
