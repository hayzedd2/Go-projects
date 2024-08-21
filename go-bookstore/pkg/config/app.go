package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func connectDB(){
	d, err := gorm.Open("mysql","azeezalhameen1@gmail.com:Hayzedd2@/go-bookstore?charset=utf8&parseTime=true&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func getDb() *gorm.DB{
	return db
}