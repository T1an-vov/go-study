package dao

import (
	"Work5/go/module"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func SqlInit() (err error) {
	db, err := gorm.Open("mysql", "root:root@/jike?charset=utf8mb4")
	if err != nil {
		return err
	}
	err = db.DB().Ping()
	if err != nil {
		return err
	}
	db.AutoMigrate(&module.User{})
	db.AutoMigrate(&module.Course{})
	db.AutoMigrate(&module.Record{})
	DB = db
	return nil
}
func Close() {
	DB.Close()
}
