package main

import (
	"time"

	"github.com/LeoCosta17/SocialMedia/internal/db"
	"github.com/LeoCosta17/SocialMedia/internal/env"
	"github.com/LeoCosta17/SocialMedia/internal/handlers"
	service "github.com/LeoCosta17/SocialMedia/internal/services"
	"github.com/LeoCosta17/SocialMedia/internal/store"
	"go.uber.org/zap"
)

func main() {

	// API configs struct
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@123!*@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetTime("DB_MAX_IDLE_TIME", time.Minute*10),
		},
	}

	// Logger init
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database configs struct
	db, err := db.New(
		cfg.dbConfig.addr,
		cfg.dbConfig.maxOpenConns,
		cfg.dbConfig.maxIdleConns,
		cfg.dbConfig.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	storage := store.NewPostgresStorage(db)
	service := service.NewService(storage)
	handler := handlers.NewHandler(service)

	// API struct
	app := &application{
		config:  cfg,
		handler: handler,
		service: service,
		storage: storage,
		logger:  logger,
	}

	r := mount(app)

	logger.Fatal(app.run(r))
}
