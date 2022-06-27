package routes

import (
	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookSToreRoutes handles routing for the API endpoints
func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")

	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
