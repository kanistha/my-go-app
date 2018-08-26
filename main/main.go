package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	."mygoapp/http"
)


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", http.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
