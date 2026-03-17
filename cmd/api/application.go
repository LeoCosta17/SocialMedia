package main

import (
	"net/http"
	"time"

	"github.com/LeoCosta17/SocialMedia/internal/handlers"
	"github.com/LeoCosta17/SocialMedia/internal/services"
	"github.com/LeoCosta17/SocialMedia/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config  config
	handler handlers.Handler
	service services.Service
	storage store.Storage
	logger  *zap.SugaredLogger
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

	app.logger.Infof("server has started at port %s", app.config.addr)

	return srv.ListenAndServe()
}
