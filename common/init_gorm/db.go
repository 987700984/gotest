package init_gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func GormMysql(MysqlDb string) *gorm.DB {
	Db, err := gorm.Open(mysql.Open(MysqlDb), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ca_", // 设置表名前缀
			SingularTable: true,  // 使用单数表名，如：user 表而不是 users
		}})
	if err != nil {
		panic("连接数据库失败")
		//logger.Write(err.Error(), "error")
	}
	if Db.Error != nil {
		panic("数据库错误")
		//logger.Write("数据库错误", "error")
	}
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return Db
}
