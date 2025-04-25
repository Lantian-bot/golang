package main

import (
	db "basic/task3/gorm/dbconn"
	"basic/task3/gorm/models"
	"fmt"
)

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

func main() {
	// 1. 连接数据库
	connection := db.DBConnection()
	fmt.Println(connection)
	//var user models.User
	var post models.Post
	post.Title = "gorm"
	post.Content = "a good language"
	post.UserID = 3
	connection.Create(&post)
	var count int64
	connection.Debug().Model(&post).Where("user_id = ?", post.UserID).Count(&count)
	fmt.Println(count)

}
