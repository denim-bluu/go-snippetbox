package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.joonkang.net/internal/models"

	_ "github.com/lib/pq"
)

type application struct {
	logger       *slog.Logger
	snippetModel *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := application{logger: logger, snippetModel: &models.SnippetModel{DB: db}}

	defer db.Close()

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
		dbSource = "postgresql://myappuser:myapppassword@localhost:5433/myappdb?sslmode=disable"
	)
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
