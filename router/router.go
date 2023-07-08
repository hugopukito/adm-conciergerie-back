package router

import (
	"adame/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/forms", service.GetForms).Methods("GET", "OPTIONS")
	router.HandleFunc("/forms", service.PostForm).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
