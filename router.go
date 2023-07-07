package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/forms", GetForms).Methods("GET", "OPTIONS")
	router.HandleFunc("/forms", PostForm).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
