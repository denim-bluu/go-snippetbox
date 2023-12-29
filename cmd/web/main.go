package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.joonkang.net/internal/models"

	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	logger        *slog.Logger
	snippetModel  *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *schema.Decoder
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	addr := flag.String("addr", ":4000", "HTTP network address")

	db, err := openDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
	}

	app := application{
		logger:        logger,
		snippetModel:  &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   schema.NewDecoder(),
	}
	app.logger.Info("Starting server on localhost", "addr", *addr)

	err = http.ListenAndServe(*addr, app.newRouter())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	dbSource := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
