package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) newRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(app.logRequest, secureHeaders)
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static")))
	router.PathPrefix("/static/").Handler(s)
	router.HandleFunc("/", app.home).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/snippet/view/{id}", app.snippetView).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", app.snippetCreate).Methods(http.MethodPost)
	router.HandleFunc("/snippet/delete/{id}", app.snippetDelete).Methods(http.MethodDelete)
	return router
}
