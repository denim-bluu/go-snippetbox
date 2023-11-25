package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	app := application{logger: slog.New(slog.NewTextHandler(os.Stdout, nil))}

	router := mux.NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static")))
	router.PathPrefix("/static/").Handler(s)
	router.HandleFunc("/", app.home).Methods(http.MethodGet)
	router.HandleFunc("/snippet/view/{id}", app.snippetView).Methods(http.MethodGet)
	router.HandleFunc("/snippet/create", app.snippetCreate).Methods(http.MethodPost)
	router.HandleFunc("/snippet/delete/{id}", app.snippetDelete).Methods(http.MethodDelete)

	app.logger.Info("Starting server on localhost", "addr", *addr)

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
