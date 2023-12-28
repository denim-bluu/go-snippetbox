package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippetModel.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "home.html", templateData{Snippets: snippets})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippetModel.Get(id)
	if err != nil {
		app.serverError(w, r, err)
	}

	app.render(w, r, http.StatusOK, "view.html", templateData{Snippet: snippet})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	fmt.Println(title)
	snippet, err := app.snippetModel.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	w.Write([]byte(fmt.Sprint(snippet)))
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
