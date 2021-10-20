package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/AdiF1/solidity/bookings/helpers"
	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/driver"
	"github.com/AdiF1/solidity/bookings/internal/forms"
	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/AdiF1/solidity/bookings/internal/render"
	"github.com/AdiF1/solidity/bookings/internal/repository"
	"github.com/AdiF1/solidity/bookings/internal/repository/dbrepo"
	"github.com/go-chi/chi"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
	DB repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
 return &Repository {
	 App: a,
	 DB: dbrepo.NewPostgresRepo(db.SQL, a),
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
	
	// send the data to the template
	render.Template(w, r, "home.page.html", &models.TemplateData {})
}
// About renders the About page 
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.Template(w, r, "about.page.html", &models.TemplateData {
		//StringMap: stringMap,
	})
}
// Contact renders the contact  page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.html", &models.TemplateData {})
}
// SearchAvailability renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.html", &models.TemplateData {})
}

type jsonResponse struct {
	OK			bool	`json:"ok"`
	Message 	string	`json:"message"`
	RoomID 		string `json:"room_id"`
	StartDate 	string `json:"start_date"`
	EndDate 	string `json:"end_date"`
}
// SearchAvailabilityJSON renders the search availability page
func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	
	sd := r.Form.Get("start")
	ed := r.Form.Get("end")

	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	available, _ := m.DB.SearchAvailabilityByDatesAndRoomID(startDate, endDate, roomID)

	resp := jsonResponse{
		OK: available,
		Message: "",
		StartDate: sd,
		EndDate: ed,
		RoomID: strconv.Itoa(roomID),
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
// PostSearchAvailability handles the form on the search availability page
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(rooms) == 0 {
		// no availability, so store a message in session
		m.App.Session.Put(r.Context(), "error", "No availability")
		// then redirect to reload page for another search
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		//return
	}

	// create a reservation, add the available dates and store in session
	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	m.App.Session.Put(r.Context(), "reservation", res)

	data := make(map[string]interface{})
	data["rooms"] = rooms
	
	// send available room info to choose-room
	render.Template(w, r, "choose-room.page.html", &models.TemplateData {
		Data: data,
	})
}
// MakeReservation renders the reservation  page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cannot get reservation from session"))
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
	}

	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.html", &models.TemplateData {
		Form: forms.New(nil),
		Data: data,
		StringMap: stringMap,
	})
}
// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get res from session"))
	}
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// update the reservation from session
	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	form := forms.New(r.PostForm)

	// error checking and validating
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")


	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// store reservation struct in session for reservation-summary template use
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// after any post, redirect to a different page than the form being posted
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}
// Esthers renders the esthers page
func (m *Repository) Esthers(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "esthers.page.html", &models.TemplateData {})
}
// Sanctuary renders the sanctuary page
func (m *Repository) Sanctuary(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "sanctuary.page.html", &models.TemplateData {})
}
// ReservationSummary renders the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	// get reservation from session and include an assert in case things go wrong
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation to display here!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// remove reservation struct from session now
	m.App.Session.Remove(r.Context(), "reservation")
	
	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData {
		Data: data,
		StringMap: stringMap,
	})
}
// ChooseRoom displays list of available rooms
func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
// BookRoom takes URL params, builds sessional variable and takes user to make-res page
func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {

	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	var res models.Reservation

	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		helpers.ServerError(w, err)
	}
	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}