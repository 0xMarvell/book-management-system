package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Marvellous-Chimaraoke/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/mysql"
)

// main handles the launching of the server
func main() {
	r := mux.NewRouter()

	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("error loading envirionment variables", dotEnvErr)
	}

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
