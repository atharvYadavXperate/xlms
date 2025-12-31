package routers

import (
	"net/http"

	"github.com/atharvYadavXperate/xlms/backend/middleware"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func SetupUserRouters(r *mux.Router) {
	r.HandleFunc("/health", healthCheck).Methods("GET")
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/register", Register).Methods("POST")
	userRouter.HandleFunc("/login", Login).Methods("POST")
	userRouter.HandleFunc("/otp", GenerateOTP).Methods("POST")
	protected := userRouter.NewRoute().Subrouter()
	protected.Use(middleware.AuthMiddleWare)
	protected.HandleFunc("/getusers", GetUsers).Methods("GET")
}

func SetupRefreshTokenRouter(r *mux.Router) {
	r.HandleFunc("/refresh", RefreshAccessToken).Methods("POST")
}

func SetupAdminRouter(r *mux.Router) {
	// r.HandleFunc("/get")
}
