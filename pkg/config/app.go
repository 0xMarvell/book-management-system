package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Connect makes API to connect with MySQL database
func Connect() {
	d, err := gorm.Open("mysql", "marvel:nolongerHUM1N_@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB gets the connected database
func GetDB() *gorm.DB {
	return db
}
