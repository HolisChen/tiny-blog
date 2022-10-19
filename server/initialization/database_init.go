package initialization

import (
	"errors"
	"github.com/HolisChen/tiny-blog/model"
	"github.com/HolisChen/tiny-blog/resources"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/**
初始化数据库
*/
func InitDatabase() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiny-blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Article{}, &model.ArticleTag{}, &model.Author{})
	if err != nil {
		return err
	}

	//数据库连接池
	sqlDb, err := db.DB()
	if err != nil {
		return errors.New("数据库连接池创建失败：" + err.Error())
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(time.Hour)

	err = resources.SetDB(db)
	if err != nil {
		return err
	}
	return nil
}
