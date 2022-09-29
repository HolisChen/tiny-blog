package main

import (
	"github.com/HolisChen/tiny-blog/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Tiny Blog",
		})
	})
	api.RegisterRouter(r)
	r.Run(":8080")

}
