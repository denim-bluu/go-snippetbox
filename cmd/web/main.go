package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	app := application{logger: slog.New(slog.NewTextHandler(os.Stdout, nil))}
	db, err := openDB()
	if err != nil {
		app.logger.Error(err.Error())
	}
	db.Ping()

	app.logger.Info("Starting server on localhost", "addr", *addr)

	err = http.ListenAndServe(*addr, app.newRouter())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB() (*sql.DB, error) {
	const (
		dbDriver = "postgres"
		dbSource = "postgresql://root:secret@localhost:5432/snippet_app?sslmode=disable"
	)
	db, err := sql.Open(dbDriver, dbSource)
	return db, err
}
