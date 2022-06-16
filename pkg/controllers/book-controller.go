package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/config"
	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/models"
	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

// GetBook sends http request to retrieve all books stored in database.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

// GetBookById retrieves book from databse using its Id.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if bookExists(bookId) == false {
		w.WriteHeader(http.StatusNotFound)

		jsonEncodeErr := json.NewEncoder(w).Encode("error: book does not exist in database")
		if jsonEncodeErr != nil {
			log.Fatal(jsonEncodeErr)
		}

		return
	}

	id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

// CreateBook creates a new book and saves it in the database.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()

	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

// UpdateBook updates the details of books stored in the database.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if bookExists(bookId) == false {
		w.WriteHeader(http.StatusNotFound)

		jsonEncodeErr := json.NewEncoder(w).Encode("error: book does not exist in database")
		if jsonEncodeErr != nil {
			log.Fatal(jsonEncodeErr)
		}

		return
	}

	id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails) // save changes to book object

	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonEncodeErr := json.NewEncoder(w).Encode("book details have been updated successfully!")
	if jsonEncodeErr != nil {
		log.Fatal(jsonEncodeErr)
	}
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

// DeleteBook deletes a book (based on the book Id) from the database.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if bookExists(bookId) == false {
		w.WriteHeader(http.StatusNotFound)

		jsonEncodeErr := json.NewEncoder(w).Encode("error: book does not exist in database")
		if jsonEncodeErr != nil {
			log.Fatal(jsonEncodeErr)
		}

		return
	}

	id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}

	bookDetails := models.DeleteBook(id)
	_, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonEncodeErr := json.NewEncoder(w).Encode("book deleted successfully!")
	if jsonEncodeErr != nil {
		log.Fatal(jsonEncodeErr)
	}
}

// bookExists checks if requested book exists in database.
func bookExists(id string) bool {
	var book models.Book
	config.GetDB().First(&book, id)
	if book.ID == 0 {
		return false
	} else {
		return true
	}
}
