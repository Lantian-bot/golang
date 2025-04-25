package main

import (
	db "basic/task3/gorm/dbconn"
	"basic/task3/gorm/models"
	"fmt"
)

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：

	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func main() {
	// 1. 连接数据库
	connection := db.DBConnection()
	var users []models.User
	// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// 定义接收查询结果的结构体
	type UserPostCommentResult struct {
		ID          uint   `gorm:"column:id"`
		UserName    string `gorm:"column:name"`
		PostTitle   string `gorm:"column:title"`
		PostContent string `gorm:"column:content"`
		Comment     string `gorm:"column:comment"`
		PostID      uint   `gorm:"column:post_id"`
	}
	var results []UserPostCommentResult

	error := connection.Debug().Model(&users).Select("users.id ,users.name, posts.title, posts.content,comments.comment,comments.post_id").
		Joins("INNER JOIN posts ON users.id = posts.user_id").Joins("INNER JOIN comments ON posts.id = comments.post_id ").
		Where("users.id = ?", 1).Find(&results)
	if error != nil {
		fmt.Println(error)
	}
	for _, r := range results {
		fmt.Println(r)
	}

	// 使用Gorm查询评论数量最多的文章信息。
	var posts []models.Post
	type PostCommentCountResult struct {
		ID           uint   `gorm:"column:id"`
		UserID       uint   `gorm:"column:user_id"`
		Title        string `gorm:"column:title"`
		Content      string `gorm:"column:content"`
		CommentCount uint   `gorm:"column:comment_count"`
	}
	var countResult []PostCommentCountResult
	error = connection.Debug().Model(&posts).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&countResult)
	if error != nil {
		fmt.Println(error)
	}
	for _, count := range countResult {
		fmt.Println(count)
	}
}
