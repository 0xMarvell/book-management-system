package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0xMarvell/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

// main handles the launching of the server
func main() {
	r := mux.NewRouter()

	// dotEnvErr := godotenv.Load()
	// if dotEnvErr != nil {
	// 	log.Fatal("error loading envirionment variables")
	// }

	routes.RegisterBookStoreRoutes(r)
	log.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
