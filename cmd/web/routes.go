package main

import (
	"net/http"
	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/handlers"
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
	mux.Get("/esthers", handlers.Repo.Esthers)
	mux.Get("/sanctuary", handlers.Repo.Sanctuary)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.Repo.SearchAvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// FileServer returns a handler that serves HTTP requests with the contents 
	// of the file system rooted at root.

	// As a special case, the returned file server redirects any request ending 
	// in "/index.html" to the same path, without the final "index.html".

	// To use the operating system's file system implementation, use http.Dir:
	fileServer := http.FileServer(http.Dir("./static/"))

	//Handle adds the route `pattern` that matches any http method to execute the `handler` http.Handler.

	// StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the 
	// request URL's Path (and RawPath if set) and invoking the handler h. StripPrefix handles a request 
	// for a path that doesn't begin with prefix by replying with an HTTP 404 not found error. The prefix 
	// must match exactly: if the prefix in the request contains escaped characters the reply is also an 
	// HTTP 404 not found error.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}