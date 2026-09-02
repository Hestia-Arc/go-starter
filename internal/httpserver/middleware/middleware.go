package middleware

import (
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func Chain(handler http.Handler, middleware ...Middleware) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler
}

func Default() []Middleware {
	return []Middleware{
		Recovery,
		RequestID,
		Logging,
		Timeout(30 * time.Second),
		SecurityHeaders,
	}
}
