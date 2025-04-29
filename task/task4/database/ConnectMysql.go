package database

import (
	"basic/task4/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func DBConnection() *gorm.DB {
	//1. 数据库参数准备
	DB_USER := "root"
	DB_PASSWORD := "Test*123456"
	DB_HOST := "192.168.45.130"
	DB_PORT := 3306
	DB_NAME := "Blog"
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, charset)
	//日志处理
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level Silent、Error、Warn、Info
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	//2. 创建数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 3.注册数据库模型
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	return db
}
