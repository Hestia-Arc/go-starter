package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hestia-Arc/go-starter/internal/app"
	"github.com/Hestia-Arc/go-starter/internal/config"
)

func main() {
	cfg := config.Load()

	application := app.New(cfg)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	err := application.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
