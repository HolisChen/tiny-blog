package model

import "time"

type ArticleTag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	ArticleId uint   `json:"article_id"`
	TagName   string `json:"tag_name" gorm:"type:varchar(10)";not null;index`
}
