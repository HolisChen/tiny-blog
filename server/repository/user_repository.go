package repository

import (
	"github.com/HolisChen/tiny-blog/model"
	"github.com/HolisChen/tiny-blog/resources"
	log "github.com/sirupsen/logrus"
)

func FindUserByLoginId(loginId string) *model.Author {
	author := &model.Author{}
	rows, err := resources.GetDB().Where("login_id = ?", loginId).First(author).Rows()
	if err != nil {
		log.Error("查询失败：", err)
		return nil
	}
	if rows.Next() {
		return author
	}
	return nil
}
