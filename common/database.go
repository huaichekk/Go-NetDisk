package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
)

func DB() *gorm.DB {
	if db == nil {
		once.Do(func() {
			var err error
			dsn := "root:320930@tcp(124.70.56.84:3306)/go-netdisk?charset=utf8mb4&parseTime=True&loc=Local"
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
		})
	}
	return db
}
