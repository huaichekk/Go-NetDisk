package models

import "gorm.io/gorm"

type UserFile struct {
	gorm.Model
	UserID       int
	FatherFileID int
	FileID       int
	Type         int
}

func init() {
	RegisterModel(&UserFile{})
}
