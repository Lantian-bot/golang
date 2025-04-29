package main

import (
	"basic/task4/database"
	"basic/task4/models"
	"fmt"
	"time"
)

func main() {
	println("hello world")
	//连接数据库
	conn := database.DBConnection()
	fmt.Println(conn)
	var posts models.Post
	m := make(map[string]interface{})
	m["content"] = "西游记介绍111"
	m["title"] = "西游记"
	m["updated_at"] = time.Now()
	result := conn.Debug().Model(&posts).Where("id = ? and user_id = ?", 1, 19).Updates(m)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

}
