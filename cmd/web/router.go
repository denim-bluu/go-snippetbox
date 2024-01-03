package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"snippetbox.joonkang.net/ui"
)

func (app *application) newRouter() *mux.Router {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.FS(ui.Files))
	WrapHandler := alice.New(app.recoverPanic, app.logRequest, secureHeaders, app.authenticate)
	protetedHandler := WrapHandler.Append(app.requireAuthentication)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	router.HandleFunc("/ping", app.ping).Methods(http.MethodGet)

	router.Handle("/static/*filepath", fileServer).Methods(http.MethodGet)
	router.Handle("/", WrapHandler.ThenFunc(app.home)).Methods(http.MethodGet, http.MethodHead)
	router.Handle("/snippet/view/{id}", WrapHandler.ThenFunc(app.snippetView)).Methods(http.MethodGet)
	router.Handle("/user/signup", WrapHandler.ThenFunc(app.userSignup)).Methods(http.MethodGet)
	router.Handle("/user/signup", WrapHandler.ThenFunc(app.userSignupPost)).Methods(http.MethodPost)
	router.Handle("/user/login", WrapHandler.ThenFunc(app.userLogin)).Methods(http.MethodGet)
	router.Handle("/user/login", WrapHandler.ThenFunc(app.userLoginPost)).Methods(http.MethodPost)

	router.Handle("/snippet/create", protetedHandler.ThenFunc(app.snippetCreate)).Methods(http.MethodGet)
	router.Handle("/snippet/create", protetedHandler.ThenFunc(app.snippetCreatePost)).Methods(http.MethodPost)
	router.Handle("/snippet/delete", protetedHandler.ThenFunc(app.snippetDelete)).Methods(http.MethodGet)
	router.Handle("/snippet/delete", protetedHandler.ThenFunc(app.snippetDeletePost)).Methods(http.MethodPost)
	router.Handle("/user/logout", protetedHandler.ThenFunc(app.userLogoutPost)).Methods(http.MethodPost)
	return router
}
