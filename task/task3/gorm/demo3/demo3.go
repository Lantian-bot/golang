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
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

func main() {
	// 1. 连接数据库
	connection := db.DBConnection()
	var comments []models.Comment
	// 删除评论 按照 post_id 删除
	// var comment models.Comment
	//result := connection.Debug().Where("post_id = ?", comment.PostID).Delete(&comment)

	// 如果不按照 post_id 删除，则先查询出 post_id
	result := connection.Debug().Where("post_id = ?", 11).Find(&comments)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, comment := range comments {
		fmt.Println(comment)
		fmt.Println(comment.ID)
		connection.Debug().Delete(&comment, comment.ID)
	}

}
