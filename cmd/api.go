package main

import (
	"net/http"
)

type application struct {
	config config
}

func (api *application) mount() http.Handler {

}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
