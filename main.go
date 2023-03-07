package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	// Get cli args
	args := os.Args[1:]
	port := "8080"

	if len(args) > 0 {
		if port != "" {
			port = args[0]
		}
	}

	listenAddr := fmt.Sprintf(":%s", port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "application/json"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", index)
		r.Get("/healthz", healthz)
	})

	log.Fatal(http.ListenAndServe(listenAddr, r))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"msg": go-hello}`)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}
