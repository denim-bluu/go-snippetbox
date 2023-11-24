package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("Snippet Viewing: %d", id)))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet Create"))
}
func snippetDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("Snippet Deleting: %d", id)))
}
