package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/models"
)

var NewBook models.Book

// GetBook sends http request to retrieve all books.
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
