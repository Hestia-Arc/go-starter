package httpserver

import (
	"net/http"

	"github.com/Hestia-Arc/go-starter/internal/httpserver/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return middleware.Chain(
		mux,
		middleware.Default()...,
	)
}
