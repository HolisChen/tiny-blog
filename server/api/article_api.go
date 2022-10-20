package api

import (
	"github.com/HolisChen/tiny-blog/model/response"
	"github.com/HolisChen/tiny-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取文章列表 godoc
// @Summary 获取文章列表
// @Schemes
// @Description 执行登录返回token
// @Tags 文章
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} string "Response"
// @Router /article/list [get]
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(service.GetArticles()))
}

func CreateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(1))
}
