package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// useCORS handle the cors config for request from the browser
func useCORS() func(http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:   []string{"*"}, //change soon
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}

	return cors.Handler(options)
}

func JSONHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// WrapMiddleware is the list of middleware that will be used for public & private routes
func WrapMiddleware() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		useCORS(),
		middleware.Logger,
		middleware.Recoverer,
		JSONHeader,
	}
}
