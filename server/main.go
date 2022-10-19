package main

import (
	"github.com/HolisChen/tiny-blog/api"
	"github.com/HolisChen/tiny-blog/docs"
	"github.com/HolisChen/tiny-blog/initialization"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func main() {
	go systemInit()
	err := newServer().Run(":8080")
	if err != nil {
		log.Panic(err)
	}
}

func newServer() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Tiny Blog",
		})
	})
	//register swagger router
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.RegisterRouter(r)
	return r
}

/**
系统初始化工作
*/
func systemInit() {
	//初始化DB
	log.Info("开始初始化数据库。。。")
	err := initialization.InitDatabase()
	if err != nil {
		log.Error("数据库初始化失败：", err)
		return
	}
	log.Info("数据库初始化完毕！")

}
