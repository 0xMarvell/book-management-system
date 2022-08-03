package config

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const projectDirName = "book-management-system" // change to relevant project name

/* loadEnv loads data from the .env file using a dynamic path.
   This helps with preventing errors while loading env files
   and makes the process of running tests in different subdirectories smooth.
*/
func loadEnv() {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Connect makes API to connect with MySQL database
func Connect() {
	loadEnv()

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
