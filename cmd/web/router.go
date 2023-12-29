package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) newRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(app.recoverPanic, app.logRequest, secureHeaders)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static")))
	router.PathPrefix("/static/").Handler(s)
	router.HandleFunc("/", app.home).Methods(http.MethodGet, http.MethodHead)
	router.HandleFunc("/snippet/view/{id}", app.snippetView).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", app.snippetCreate).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", app.snippetCreatePost).Methods(http.MethodPost)
	router.HandleFunc("/snippet/remove", app.snippetRemove).Methods(http.MethodGet)
	router.HandleFunc("/snippet/remove", app.snippetRemoveDelete).Methods(http.MethodPost)
	router.HandleFunc("/user/signup", app.userSignup).Methods(http.MethodGet)
	router.HandleFunc("/user/signup", app.userSignupPost).Methods(http.MethodPost)
	router.HandleFunc("/user/login", app.userLogin).Methods(http.MethodGet)
	router.HandleFunc("/user/login", app.userLoginPost).Methods(http.MethodPost)
	router.HandleFunc("/user/logout", app.userLogoutPost).Methods(http.MethodPost)

	return router
}
