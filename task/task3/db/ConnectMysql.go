package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// go get github.com/go-sql-driver/mysql
// go get github.com/jmoiron/sqlx

var (
	//1. 数据库参数准备
	DB_USER     string = "root"
	DB_PASSWORD string = "Test*123456"
	DB_HOST     string = "192.168.45.130"
	DB_PORT     int    = 3306
	DB_NAME     string = "web3"
	charset     string = "utf8mb4"
)

// 连接数据库
func ConnectMySQL() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, charset)
	// 2. 创建数据库连接
	db, err := sqlx.Open("mysql", dsn)
	fmt.Println(err)
	ping(db)
	return db
}

// 测试数据库连接是否成功
func ping(DB *sqlx.DB) {
	err := DB.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据库连接成功")
	}
}
