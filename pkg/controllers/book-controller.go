package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/0xMarvell/book-management-system/pkg/config"
	"github.com/0xMarvell/book-management-system/pkg/models"
	"github.com/0xMarvell/book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

// HomePage displays simple message on the home page.
func HomePage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"success": true,
		"message": "Hello!, try sending HTTP requests to the endpoints specified in the documentation :)",
	}

	displayMessage, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(displayMessage)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}

// GetBook sends http request to retrieve all books stored in database.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.MarshalIndent(newBooks, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}

// GetBookById retrieves book from databse using its Id.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if !bookExists(bookId) {
		displayErrorMessage(w, map[string]interface{}{
			"success": false,
			"message": "book does not exist in database",
		})

		return
	}

	id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		log.Fatalf("error while parsing: %v", err1)
	}

	bookDetails, _ := models.GetBookById(id)
	res, err := json.MarshalIndent(bookDetails, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}

// CreateBook creates a new book and saves it in the database.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()

	res, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}

// UpdateBook updates the details of books stored in the database.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if !bookExists(bookId) {
		displayErrorMessage(w, map[string]interface{}{
			"success": false,
			"message": "book does not exist in database",
		})

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

	res, err := json.MarshalIndent(bookDetails, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonEncodeErr := json.NewEncoder(w).Encode("book details have been updated successfully!")
	if jsonEncodeErr != nil {
		log.Fatalf("error encoding JSON: %v", jsonEncodeErr)
	}

	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}

// DeleteBook deletes a book (based on the book Id) from the database.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	if !bookExists(bookId) {
		displayErrorMessage(w, map[string]interface{}{
			"success": false,
			"message": "book does not exist in database",
		})

		return
	}

	id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}

	bookDetails := models.DeleteBook(id)
	res, err := json.MarshalIndent(bookDetails, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonEncodeErr := json.NewEncoder(w).Encode("book deleted successfully!")
	if jsonEncodeErr != nil {
		log.Fatalf("error encoding JSON: %v", jsonEncodeErr)
	}

	_, writeErr := w.Write(res)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
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

// displayErrorMessage displays an error message as json after sending a http request
func displayErrorMessage(w http.ResponseWriter, data map[string]interface{}) {
	displayErrorMessage, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("error marshalling JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, writeErr := w.Write(displayErrorMessage)
	if writeErr != nil {
		log.Fatalf("could not write data: %v", writeErr)
	}
}
