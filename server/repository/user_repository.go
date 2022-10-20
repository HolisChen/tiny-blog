package repository

import (
	"github.com/HolisChen/tiny-blog/model"
	"github.com/HolisChen/tiny-blog/resources"
	log "github.com/sirupsen/logrus"
	"time"
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

func FindByMail(mail string) *model.Author {
	author := &model.Author{}
	rows, err := resources.GetDB().Where("mail_address = ?", mail).First(author).Rows()
	if err != nil {
		log.Error("查询失败：", err)
		return nil
	}
	if rows.Next() {
		return author
	}
	return nil
}

func Insert(author *model.Author) error {
	_, err := resources.GetDB().Create(author).Rows()
	if err != nil {
		return err
	}
	return nil
}

func UpdateLastLoginTime(userId uint) {
	resources.GetDB().Model(&model.Author{}).Where("id = ?", userId).Update("last_login_time", time.Now())
}
