package dao

import (
	"example.com/v2/config"
	"example.com/v2/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MysqlDb), &gorm.Config{})
	if err != nil {

		logger.Write(err.Error(), "error")
	}
	if Db.Error != nil {

		logger.Write("数据库错误", "error")
	}
	sqlDB, err := Db.DB()
	if err != nil {
		logger.Write(err.Error(), "error")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logger.Write("787545", "error")
}
