package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null;comment:文件名"`              // [1,5,6](@ref)
	Size   int64  `gorm:"not null;default:0;comment:文件大小"`                     // [5,7](@ref)
	Hash   string `gorm:"type:varchar(255);uniqueIndex;not null;comment:文件哈希"` // [1,3,6](@ref)
	Status int    `gorm:"index;default:1;comment:状态(0未启用/1正常)"`                // [3,5](@ref)
}

func init() {
	RegisterModel(&File{})
}
