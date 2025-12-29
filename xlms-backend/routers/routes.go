package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.HandleFunc("/users/register", Register).Methods("POST")
	r.HandleFunc("/users/login", Login).Methods("POST")
	r.HandleFunc("/users/otp", GenerateOTP).Methods("POST")
	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
