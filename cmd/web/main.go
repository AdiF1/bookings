package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiF1/solidity/bookings/pkg/config"
	"github.com/AdiF1/solidity/bookings/pkg/handlers"
	"github.com/AdiF1/solidity/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"

)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// store in app-side variable in config.go
	app.Session = session

	// create a template cache
	tc, err := render.CreateTemplateCache()
	//fmt.Fprintln(tc)

	if err != nil {
		log.Fatal("error creating template cache")
	}

	app.TemplateCache = tc

	// SET TRUE FOR PRODUCTION
	app.UseCache = false

	// calls a func to create a new repository
	repo := handlers.NewRepo(&app)

	// calls a func to set the repository for the handlers
	handlers.NewHandlers(repo)

	// calls a func to set the repository for the handlers
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	srv := &http.Server{

		Addr : portNumber,
		Handler : routes(&app),
	}

	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler 
	// to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	err = srv.ListenAndServe()
	log.Fatal(err)
}





