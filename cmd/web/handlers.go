package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"snippetbox.joonkang.net/internal/models"
	"snippetbox.joonkang.net/internal/validator"
)

type snippetCreateForm struct {
	Title               string `schema:"title,required"`
	Content             string `schema:"content,required"`
	Expires             int    `schema:"expires,required"`
	validator.Validator `schema:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippetModel.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.html", data)
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
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	session, err := app.cookieStore.Get(r, "session")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fm := session.Flashes("create-message")
	if fm != nil {
		data.Flash = fm[0].(string)
		session.Save(r, w)
	}

	app.render(w, r, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		Expires: 365,
	}
	app.render(w, r, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.Check(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.Check(validator.MaxStringLength(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.Check(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.Check(validator.PermitteValues(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	snippet, err := app.snippetModel.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	session, err := app.cookieStore.Get(r, "session")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	session.AddFlash("Snippet successfully created!", "create-message")
	session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippet.ID), http.StatusSeeOther)
}

func (app *application) snippetRemove(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	ids, err := app.snippetModel.GetIDs()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data.IDs = ids
	app.render(w, r, http.StatusOK, "remove.html", data)
}

func (app *application) snippetRemoveDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PostForm.Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	err = app.snippetModel.Delete(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}
