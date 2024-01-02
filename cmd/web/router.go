package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
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
	router.HandleFunc("/user/signup", app.userSignup).Methods(http.MethodGet)
	router.HandleFunc("/user/signup", app.userSignupPost).Methods(http.MethodPost)
	router.HandleFunc("/user/login", app.userLogin).Methods(http.MethodGet)
	router.HandleFunc("/user/login", app.userLoginPost).Methods(http.MethodPost)

	protected := alice.New(app.requireAuthentication)
	router.Handle("/snippet/create", protected.ThenFunc(app.snippetCreate)).Methods(http.MethodGet)
	router.Handle("/snippet/create", protected.ThenFunc(app.snippetCreatePost)).Methods(http.MethodPost)
	router.Handle("/snippet/delete", protected.ThenFunc(app.snippetDelete)).Methods(http.MethodGet)
	router.Handle("/snippet/delete", protected.ThenFunc(app.snippetDeletePost)).Methods(http.MethodPost)
	router.Handle("/user/logout", protected.ThenFunc(app.userLogoutPost)).Methods(http.MethodPost)
	return router
}
