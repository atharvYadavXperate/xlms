package main

import (
	"log"
	"net/http"

	"github.com/atharvYadavXperate/xlms/backend/routers"
	db "github.com/atharvYadavXperate/xlms/database"
	"github.com/gorilla/handlers"
)

func main() {
	db.ConnectDb()

	defer db.CloseConnection()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	router := routers.SetupRouters()
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
