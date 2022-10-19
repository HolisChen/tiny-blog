package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1 := apiGroup.Group("/v1")
	articleRouter := v1.Group("/article")
	/*Article API start*/
	articleRouter.GET("/list", GetArticles)
	articleRouter.POST("/create", CreateArticle)

	/*用户（作者）相关API*/
	userRouter := v1.Group("/user")
	userRouter.POST("/login", Login)
	userRouter.POST("/logout", Logout)
	userRouter.POST("/register", Register)
	userRouter.POST("/resetPwd", ResetPwd)

}
