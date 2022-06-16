package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect makes API to connect with MySQL database
func Connect() {
	// Load env vars
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("error loading envirionment variables")
	}

	dbUsername := os.Getenv("DBUSERNAME")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB gets the connected database
func GetDB() *gorm.DB {
	return db
}
