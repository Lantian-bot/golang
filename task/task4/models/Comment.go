package models

import "gorm.io/gorm"

// comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}
