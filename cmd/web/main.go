package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	addr := ":4000"

	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/snippet/view/{id}", snippetView).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", snippetCreate).Methods(http.MethodPost)
	router.HandleFunc("/snippet/delete/{id}", snippetDelete).Methods(http.MethodDelete)

	log.Printf("Starting server on localhost%s/", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
