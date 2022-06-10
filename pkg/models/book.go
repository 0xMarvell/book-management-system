package models

import (
	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book is the model for each Book to be stored in the databse
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init initializes the database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
