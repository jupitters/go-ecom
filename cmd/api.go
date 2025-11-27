package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

func (api *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	return r
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
