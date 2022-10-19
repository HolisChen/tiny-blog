package model

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	gorm.Model
	Name          string    `json:"name" gorm:"type:varchar(50);not null;"`
	LoginId       string    `json:"loginId" gorm:"type:varchar(15);not null;uniqueIndex"`
	Password      string    `gorm:"type:varchar(50);not null"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	Level         string    `json:"level" gorm:"type:varchar(10)"`
	Status        int32     `json:"status" gorm:"type:tinyint(4);default:0;not null"` //0:正常 1:locked
	MailAddress   string    `json:"mailAddress" gorm:"type:varchar(50)"`
}
