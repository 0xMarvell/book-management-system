package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect makes API to connect with MySQL database
func Connect() {
	dsn := "marvel:nolongerHUM1N_@tcp(localhost)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//d, err := gorm.Open("mysql", "marvel:nolongerHUM1N_@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB gets the connected database
func GetDB() *gorm.DB {
	return db
}
