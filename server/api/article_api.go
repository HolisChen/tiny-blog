package api

import (
	"github.com/HolisChen/tiny-blog/model/response"
	"github.com/HolisChen/tiny-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(service.GetArticles()))
}

func CreateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(1))
}
