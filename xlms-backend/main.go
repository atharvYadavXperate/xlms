package main

import (
	"log"
	"net/http"

	"github.com/atharvYadavXperate/xlms/backend/routers"
	db "github.com/atharvYadavXperate/xlms/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDb()
	defer db.CloseConnection()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:5173",
			"https://nn5x9hqz-8080.inc1.devtunnels.ms/",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router := mux.NewRouter()

	routers.SetupUserRouters(router)
	routers.SetupAdminRouter(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
