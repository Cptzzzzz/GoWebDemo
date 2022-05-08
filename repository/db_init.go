package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		USERNAME, PASSWORD, IP, PORT, DATABASE)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		db.AutoMigrate(&Post{})
		db.AutoMigrate(&Topic{})
		db.AutoMigrate(&Id{})
	}
	return err
}
