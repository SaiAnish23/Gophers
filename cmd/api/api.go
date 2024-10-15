package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SaiAnish23/Gophers/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	// mux := http.NewServeMux()
	// // mux.HandleFunc("/", app.home)
	// mux.HandleFunc("GET /v1/health", app.healthcheck)
	// return mux
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// r.Get("/v1/health", app.healthcheck)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthcheck)
	})
	return r

}

func (app *application) run(mux *chi.Mux) error {
	// Start the server

	// mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on %s", srv.Addr)

	return srv.ListenAndServe()
}
