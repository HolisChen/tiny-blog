package model

import "gorm.io/gorm"

type Article struct {
	Title       string       `json:"title" gorm:"type:varchar(100);not null;"`
	SubTitle    string       `json:"subTitle" gorm:"type:varchar(200)"`
	Description string       `json:"description" gorm:"type:varchar(500)"`
	Content     string       `json:"content" gorm:"type:text"`
	AuthorId    int32        `json:"authorId"  gorm:"type int(11)"`
	Tags        []ArticleTag `json:"tags"`
	gorm.Model
}
