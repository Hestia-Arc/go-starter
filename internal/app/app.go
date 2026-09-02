package app

import (
	"context"
	"log"
	"time"

	"github.com/Hestia-Arc/go-starter/internal/config"
	"github.com/Hestia-Arc/go-starter/internal/httpserver"
)

type App struct {
	config config.Config
	server *httpserver.Server
}

func New(cfg config.Config) *App {
	router := httpserver.NewRouter()

	server := httpserver.New(
		":"+cfg.HTTP.Port,
		router,
	)

	return &App{
		config: cfg,
		server: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	serverErr := make(chan error, 1)

	go func() {
		log.Printf("http server listening on %s", ":"+a.config.HTTP.Port)

		serverErr <- a.server.Start()
	}()

	select {
	case err := <-serverErr:
		return err

	case <-ctx.Done():
		log.Println("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	return a.server.Shutdown(shutdownCtx)
}
