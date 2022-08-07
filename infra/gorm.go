package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 获取Gorm DB连接句柄
func NewDB() *gorm.DB {
	dsn := "root:zxcvbnm123@tcp(127.0.0.1:3306)/finance_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

// todo 迁移表function