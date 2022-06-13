package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const PORT = ":8080"

// main handles the launching of the server
func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(PORT, r))
}
