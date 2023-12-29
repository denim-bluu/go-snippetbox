package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"snippetbox.joonkang.net/internal/models"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	logger        *slog.Logger
	snippetModel  *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *schema.Decoder
	cookieStore   *sessions.CookieStore
}

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	store.Options.MaxAge = int(12 * time.Hour)
	store.Options.HttpOnly = true
	store.Options.Secure = true
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
		cookieStore:   store,
	}
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.newRouter(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app.logger.Info("Starting server on localhost", "addr", *addr)

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	logger.Error(err.Error())
	os.Exit(1)
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
