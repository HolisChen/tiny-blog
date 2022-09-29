package api

import (
	"github.com/HolisChen/tiny-blog/model/response"
	"github.com/HolisChen/tiny-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleApi struct {
}

func GetArticles(c *gin.Context) {
	articles := service.ArticleServiceApp.GetArticles()
	c.JSON(http.StatusOK, response.Success(articles))
}

func CreateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(1))
}
