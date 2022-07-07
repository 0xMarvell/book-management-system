package config

import (
	"fmt"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect makes API to connect with MySQL database
func Connect() {
	//Load env vars
	// dotEnvErr := godotenv.Load()
	// if dotEnvErr != nil {
	// 	log.Fatalf("error loading envirionment variables: %v", dotEnvErr)
	// }

	dbUsername := os.Getenv("DBUSERNAME")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")
	tcp := os.Getenv("TCP")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, tcp, dbName)
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
