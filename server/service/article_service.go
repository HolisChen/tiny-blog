package service

import (
	"github.com/HolisChen/tiny-blog/model"
)

func GetArticles() []model.Article {
	articles := make([]model.Article, 0)
	articles = append(articles, model.Article{
		Title:       "Golang进阶",
		SubTitle:    "golang中级进阶教程",
		Description: "这个适合Go语言学习",
		Content:     "这个是文章内容",
		AuthorId:    1,
	})
	return articles
}
