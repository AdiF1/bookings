package main

import (
	"net/http"

	"github.com/AdiF1/solidity/bookings/pkg/config"
	"github.com/AdiF1/solidity/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes (app *config.AppConfig) http.Handler {

	// NewRouter returns a new Mux object that implements the Router interface.
	mux := chi.NewRouter()
	// Recoverer is a middleware that recovers from panics, logs the panic 
	// (and a backtrace), and returns a HTTP 500 (Internal Server Error) status 
	// if possible. Recoverer prints a request ID if one is provided.
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}