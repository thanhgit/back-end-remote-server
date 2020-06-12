package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB

func GetInstance() *gorm.DB {
	if instance == nil {
		db, err := gorm.Open("mysql", "rsuser:123456@/remoteserver?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			println(err.Error())
			defer db.Close()
		} else {
			println("Connect success to Mysql server")
			instance = db
			instance.AutoMigrate(&Server{},&Website{})
		}

	}

	return instance
}