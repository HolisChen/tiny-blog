package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1 := apiGroup.Group("/v1")
	/*Article API start*/
	v1.GET("/articles", GetArticles)
	v1.POST("/article", CreateArticle)
}
