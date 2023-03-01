package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herizal95/golang-jwt-mux/config"
)

func main() {

	config.ConnectDatabase()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
