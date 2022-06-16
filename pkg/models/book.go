package models

import (
	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/config"
	"gorm.io/gorm"
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

// CreateBook creates a new book object in the database.
func (b *Book) CreateBook() *Book {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all books stored in the database.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById searches for books using their ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

// DeleteBook deletes book from database.
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)

	return book
}
