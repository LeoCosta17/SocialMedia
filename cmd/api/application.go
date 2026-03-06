package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LeoCosta17/SocialMedia/internal/handlers"
	"github.com/LeoCosta17/SocialMedia/internal/services"
	"github.com/LeoCosta17/SocialMedia/internal/store"
)

type application struct {
	config  config
	handler handlers.Handler
	service services.Service
	storage store.Storage
}

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func (app *application) run(r http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
