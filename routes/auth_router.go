package routes

import (
	"github.com/gorilla/mux"
	"github.com/herizal95/golang-jwt-mux/controller/authcontroller"
)

func AuthenticationRoutes(r *mux.Router) {

	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
}
