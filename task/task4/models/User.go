package models

import "gorm.io/gorm"

// users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}
