package models

import (
	"fmt"
	"gorm.io/gorm"
)

// Comment 模型
type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Comment string
	PostID  uint // 外键
}

//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("after Delete  ", comment)
	if comment.PostID == 0 {
		return nil
	}
	var count int64
	tx.Debug().Model(&comment).Where("post_id = ?", comment.PostID).Count(&count)
	fmt.Println(count)
	if count == 0 {
		return tx.Debug().Model(&Post{}).Where("id = ?", comment.PostID).Update("status_comment", 0).Error
	}
	return
}
