package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herizal95/golang-jwt-mux/config"
	"github.com/herizal95/golang-jwt-mux/routes"
)

func main() {

	config.ConnectDatabase()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthenticationRoutes(router)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
