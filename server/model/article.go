package model

import "github.com/HolisChen/tiny-blog/model/common"

type Article struct {
	Title       string   `json:"title"`
	SubTitle    string   `json:"subTitle"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	AuthorId    int32    `json:"authorId"`
	common.TimeInfo
}
