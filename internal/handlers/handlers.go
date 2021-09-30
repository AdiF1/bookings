package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/AdiF1/solidity/bookings/internal/render"
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

// Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// get and store the request IP in session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Home Page!"

	// send the data to the template
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData {
		StringMap: stringMap,
	})
}

// About renders the About page 
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	t := time.Now()
	stringMap["test"] = t.String()

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData {
		StringMap: stringMap,
	})
}
// Contact renders the contact  page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData {})
}
// SearchAvailability renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData {})
}

type jsonResponse struct {
	OK	bool	`json:"ok"`
	Message string	`json:"message"`
}

// SearchAvailabilityJSON renders the search availability page
func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")

if err != nil {
	log.Println(err)
}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// PostSearchAvailability handles the form on the search availability page
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Handling form post: start date is %s and end date is %s", start, end)))
}
// MakeReservation renders the reservation  page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData {})
}
// Esthers renders the esthers page
func (m *Repository) Esthers(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "esthers.page.html", &models.TemplateData {})
}
// Sanctuary renders the sanctuary page
func (m *Repository) Sanctuary(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "sanctuary.page.html", &models.TemplateData {})
}

