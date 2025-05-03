package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null;uniqueIndex"` // 用户名唯一索引
	Password string `gorm:"size:255;not null"`             // 密码不设索引
	IsVIP    bool   `gorm:"default:false"`
}

func init() {
	RegisterModel(&User{})
}
