package handlers

import (
	"net/http"
	"time"

	"github.com/AdiF1/solidity/bookings/pkg/config"
	"github.com/AdiF1/solidity/bookings/pkg/models"
	"github.com/AdiF1/solidity/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
 return &Repository {
	 App: a,
 }
}

// NewHandlers sets the repository for the handlers
func NewHandlers (r *Repository) { 
	Repo = r
}

// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// A Request represents an HTTP request received by a server or to be sent by a client.

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// get and store the request IP in session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Home Page!"

	// send the data to the template
	render.RenderTemplate(w, "home.page.html", &models.TemplateData {
		StringMap: stringMap,
	})
}

// About is the About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	t := time.Now()
	stringMap["test"] = t.String()

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData {
		StringMap: stringMap,
	})
}

