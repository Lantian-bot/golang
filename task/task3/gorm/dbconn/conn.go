package db

import (
	"basic/task3/gorm/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	//1. 数据库参数准备
	DB_USER := "root"
	DB_PASSWORD := "Test*123456"
	DB_HOST := "192.168.45.130"
	DB_PORT := 3306
	DB_NAME := "web3"
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 3.注册数据库模型
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	return db
}
