package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	w.Write([]byte(fmt.Sprintf("Snippet Viewing: %d", id)))
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet Create"))
}
func (app *application) snippetDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	w.Write([]byte(fmt.Sprintf("Snippet Deleting: %d", id)))
}
