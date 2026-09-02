package app

import (
	"context"
	"log"

	"github.com/Hestia-Arc/go-starter/internal/config"
)

type App struct {
	config config.Config
}

func New(cfg config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	log.Printf(
		"starting application name=%s environment=%s",
		a.config.App.Name,
		a.config.App.Environment,
	)

	<-ctx.Done()

	log.Println("shutting down application")

	return nil
}
