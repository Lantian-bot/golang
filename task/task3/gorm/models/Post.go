package models

import (
	"fmt"
	"gorm.io/gorm"
)

// Post 模型
type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Content       string
	StatusComment int       // 0: no comment, 1: has comment
	UserID        uint      // 外键
	Comments      []Comment `gorm:"foreignKey:PostID"` // 一对多关系
}

// 结构体方法

func (post *Post) BeforCreate(tx *gorm.DB) (err error) {
	fmt.Println("before create ", post)
	return
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("after create... ", post)
	if post.UserID == 0 {
		return nil
	}
	// 查询用户的文章数量
	var count int64
	tx.Debug().Model(&post).Where("user_id = ?", post.UserID).Count(&count)
	return tx.Model(&User{}).Where("id = ?", post.UserID).Update("count", count).Error
}
