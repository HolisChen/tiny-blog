package resources

import (
	"errors"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func SetDB(d *gorm.DB) error {
	if db != nil {
		return errors.New("数据库连接已初始化，不能再次初始化")
	}
	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}
