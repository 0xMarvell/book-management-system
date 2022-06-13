package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// GetBook sends http request to retrieve all books stored in database.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById retrieves book from databse using its Id.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err1 := strconv.ParseInt(bookId, 0, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(Id)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
