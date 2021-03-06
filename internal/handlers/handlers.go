package handlers

import (

	//"crypto/ecdsa"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AdiF1/solidity/bookings/contracts"
	"github.com/AdiF1/solidity/bookings/ether"
	"github.com/AdiF1/solidity/bookings/helpers"
	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/AdiF1/solidity/bookings/internal/driver"
	"github.com/AdiF1/solidity/bookings/internal/forms"
	"github.com/AdiF1/solidity/bookings/internal/models"
	"github.com/AdiF1/solidity/bookings/internal/render"
	"github.com/AdiF1/solidity/bookings/internal/repository"
	"github.com/AdiF1/solidity/bookings/internal/repository/dbrepo"
	"github.com/go-chi/chi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/crypto"
)

// Repo the repository used by the handlers
var Repo *Repository
// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}
// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}
// NewTestRepo creates a new repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingsRepo(a),
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}
// About is the about page handler
func (m *Repository) Smallchange(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.Template(w, r, "smallchange.page.html", &models.TemplateData{})
}
// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't find room!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.FmtStartDate = res.StartDate.Format("2006-01-02")
	res.FmtEndDate = res.EndDate.Format("2006-01-02")
	res.RoomName = room.RoomName

	token := ether.ContractReadERC20()

	data := make(map[string]interface{})
	data["reservation"] = res
	data["token"] = token

	m.App.Session.Put(r.Context(), "reservation", res)



	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
	})
}
// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	log.Println("2")
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
		StartDate: res.StartDate,
		EndDate:   res.EndDate,
		RoomID:    res.RoomID,
		RoomName:  res.Room.RoomName,
		FmtStartDate: res.StartDate.Format("2006-01-02"),
		FmtEndDate: res.EndDate.Format("2006-01-02"),
		Room: res.Room,
	}

	form := forms.New(r.PostForm)

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
		m.App.Session.Put(r.Context(), "error", "can't insert reservation into database!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     res.StartDate,
		EndDate:       res.EndDate,
		RoomID:        res.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert room restriction!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// send notifications - first to guest
	htmlMessage := fmt.Sprintf(`
	
	<strong>Reservation Confirmation</strong><br><br>
		Hi %s,<br><br>
		This is to confirm your reservation for %s from %s to %s.`, 
		reservation.FirstName, reservation.RoomName, reservation.StartDate.Format("2006-01-02"), 
		reservation.EndDate.Format("2006-01-02"))

	msg := models.MailData{
		To:      reservation.Email,
		From:    "info@smallchange.ca",
		Subject: "Your Reservation Confirmation",
		Content: htmlMessage,
		Template: "basic.html",
	}
	m.App.MailChan <- msg

	// then to property owner
	htmlMessage = fmt.Sprintf(`
	
	<strong>Booking Confirmation<strong><br><br>
		This is to confirm a reservation for %s from %s to %s, made by %s %s.`, 
		reservation.RoomName, reservation.StartDate.Format("2006-01-02"), 
		reservation.EndDate.Format("2006-01-02"), reservation.FirstName, reservation.LastName)

	msg = models.MailData{
		To:      "adifletcher@mac.com",
		From:    "info@smallchange.ca",
		Subject: "Booking Confirmation",
		Content: htmlMessage,
	}
	m.App.MailChan <- msg

	// store the reservation in the session
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// redirect to summary page
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}
// Esthers renders the room page
func (m *Repository) Esthers(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "esthers.page.html", &models.TemplateData{})
}
// Sanctuary renders the room page
func (m *Repository) Sanctuary(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "sanctuary.page.html", &models.TemplateData{})
}
// SearchAvailability renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.html", &models.TemplateData{})
}
// PostSearchAvailability renders the search availability page
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse start date!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse end date!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't get availability for rooms")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if len(rooms) == 0 {
		// no availability - show banner!
		m.App.Session.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.html", &models.TemplateData{
		Data: data,
	})
}

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) SearchAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// need to parse request body
	err := r.ParseForm()
	if err != nil {
		// can't parse form, so return appropriate json
		resp := jsonResponse{
			OK:      false,
			Message: "Internal server error",
		}

		out, _ := json.MarshalIndent(resp, "", "     ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))

	available, err := m.DB.SearchAvailabilityByDatesAndRoomID(startDate, endDate, roomID)
	if err != nil {
		// got a database error, so return appropriate json
		resp := jsonResponse{
			OK:      false,
			Message: "Error querying database",
		}

		out, _ := json.MarshalIndent(resp, "", "     ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}
	resp := jsonResponse{
		OK:        available,
		Message:   "",
		StartDate: sd,
		EndDate:   ed,
		RoomID:    strconv.Itoa(roomID),
	}

	// I removed the error check, since we handle all aspects of
	// the json right here
	out, _ := json.MarshalIndent(resp, "", "     ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
// Contact renders the search availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.html", &models.TemplateData{})
}
// ReservationSummary displays the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data:      data,
	})
}
// ChooseRoom displays list of available rooms
func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	
	// used to have next 6 lines
	//roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	//if err != nil {
	//	log.Println(err)
	//	m.App.Session.Put(r.Context(), "error", "missing url parameter")
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	// changed to this, so we can test it more easily
	// split the URL up by /, and grab the 3rd element
	exploded := strings.Split(r.RequestURI, "/")
	roomID, err := strconv.Atoi(exploded[2])
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "missing url parameter")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Can't get room from db!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.RoomID = roomID
	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
// BookRoom takes URL parameters, builds a sessional variable, and takes user to make res screen
func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	var res models.Reservation

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Can't get room from db!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Room.RoomName = room.RoomName
	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
// ShowLogin shows the login screen
func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostShowLogin handles logging the user in
func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request) {

	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")


	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")
	if !form.Valid() {
		log.Println("invalid login form")
		render.Template(w, r, "login.page.html", &models.TemplateData{
			Form: form,
		})
		return

		// TODO  take user back to page
	}

	id, _, err := m.DB.Authenticate(email, password)
	if err != nil {
		log.Println(err)
		m.App.Session.Put(r.Context(), "error", "Invalid login credentials")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}
	
	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "Logged in successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout logs out a user
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

// AdminDashboard shows restricted dashboard to logged in (admin) users
func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin-dashboard.page.html", &models.TemplateData{})
}

// AdminAllNewReservations shows admin page to display new reservations
func (m *Repository) AdminNewReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := m.DB.AllNewReservations()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["reservations"] = reservations

	render.Template(w, r, "admin-new-reservations.page.html", &models.TemplateData{
		Data: data,
	})
}

// AdminAllReservations shows admin page to list all reservations
func (m *Repository) AdminAllReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := m.DB.AllReservations()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["reservations"] = reservations

	render.Template(w, r, "admin-all-reservations.page.html", &models.TemplateData{
		Data: data,
	})
}

// AdminShowReservation shows a reservation in the admin tool
func (m *Repository) AdminShowReservation(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(exploded[4])
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	src := exploded[3]

	stringMap := make(map[string]string)
	stringMap["src"] = src

	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")

	stringMap["month"] = month
	stringMap["year"] = year

	// get reservation from database
	res, err := m.DB.GetReservationByID(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = res


	render.Template(w, r, "admin-show-reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
		Form: forms.New(nil),
	})

}

// AdminPostShowReservation shows a reservation in the admin tool
func (m *Repository) AdminPostShowReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	exploded := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(exploded[4])
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	src := exploded[3]
	stringMap := make(map[string]string)
	stringMap["src"] = src

	// get reservation from database
	res, err := m.DB.GetReservationByID(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	
	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.Form.Get("last_name")
	res.Phone = r.Form.Get("phone")
	res.Email = r.Form.Get("email")

	err = m.DB.UpdateReservation(res)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	month := r.Form.Get("month")
	year := r.Form.Get("year")

	m.App.Session.Put(r.Context(), "flash", "Changes saved")
	if year == "" {
			http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-calendar?y=%s&m=%s", year, month), http.StatusSeeOther)

	}
}
 
// AdminProcessReservation marks a reservation as processed
func (m *Repository) AdminProcessReservation(w http.ResponseWriter, r *http.Request) {
	// access id param in URL and convert to string
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	src := chi.URLParam(r, "src")
	// call the update function and change processed to 1
	err := m.DB.UpdateProcessedForReservation(id, 1)
	if err != nil {
		log.Println(err)
	}

	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")

	// give some feedback
	m.App.Session.Put(r.Context(), "flash", "Reservation marked as processed")

	if year == "" {
		// redirect to the correct source page
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	} else {
		// redirect to the correct calendar year/month page
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-calendar?y=%s&m=%s", year, month), http.StatusSeeOther)
	}
	
}

// AdminDeleteReservation deletes a reservation as processed
func (m *Repository) AdminDeleteReservation(w http.ResponseWriter, r *http.Request) {
	// access id param in URL and convert to string
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	src := chi.URLParam(r, "src")
	// call the update function and change processed to 1
	err := m.DB.DeleteReservation(id)
	if err != nil {
		log.Println(err)
	}

	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")

	// give some feedback
	m.App.Session.Put(r.Context(), "flash", "Reservation deleted")
	if year == "" {
		// redirect to the correct source page
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	} else {
		// redirect to the correct calendar year/month page
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-calendar?y=%s&m=%s", year, month), http.StatusSeeOther)
	}
	
}

// AdminCalendar shows admin pages to display a calendar
func (m *Repository) AdminReservationsCalendar(w http.ResponseWriter, r *http.Request) {
	// assume no month or year specified
	now := time.Now()

	if r.URL.Query().Get("y") != "" {
		year, err := strconv.Atoi(r.URL.Query().Get("y"))
		if err != nil {
			helpers.ServerError(w, err)
		return
		}
		month, err := strconv.Atoi(r.URL.Query().Get("m"))
		if err != nil {
			helpers.ServerError(w, err)
		return
		}
		now = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}

	data := make(map[string]interface{})
	data["now"] = now

	next := now.AddDate(0,1,0)
	last := now.AddDate(0, -1, 0)

	nextMonth := next.Format("01")
	nextMonthYear := next.Format("2006")
	lastMonth := last.Format("01")
	lastMonthYear := last.Format("2006")

	stringMap := make(map[string]string)
	stringMap["next_month"] = nextMonth
	stringMap["next_month_year"] = nextMonthYear
	stringMap["last_month"] = lastMonth
	stringMap["last_month_year"] = lastMonthYear

	stringMap["this_month"] = now.Format("01")
	stringMap["this_month_year"] = now.Format("2006")

	// get first and last of the month
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	intMap := make(map[string]int)
	intMap["days_in_month"] = lastOfMonth.Day()

	rooms, err := m.DB.AllRooms()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data["rooms"] = rooms

	for _, x:= range rooms {

		// create maps and init checkbox table cell data
		reservationMap := make(map[string]int)
		blockMap := make(map[string]int)

		for d := firstOfMonth; !d.After(lastOfMonth); d = d.AddDate(0, 0, 1) {
			reservationMap[d.Format("2006 01 2")] = 0
			blockMap[d.Format("2006 01 2")] = 0
		}

		// get all restrictions for the current room
		restrictions, err := m.DB.GetRestrictionsForRoomByDate(x.ID, firstOfMonth, lastOfMonth)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		for _, y := range restrictions {
			if y.ReservationID > 0 {
				// it's a reservation. taking a day off the EndDate ensures that just the overnight
				//  date shows on the calendar: it's a single night booking from 1500 - 1100 hrs.
				for date := y.StartDate; !date.After(y.EndDate.AddDate(0, 0, -1)); date = date.AddDate(0, 0, 1) {
					reservationMap[date.Format("2006-01-2")] = y.ReservationID
				}

			} else {
				// it's a block
				blockMap[y.StartDate.Format("2006-01-2")] = y.ID
			}
		}

		data[fmt.Sprintf("reservation_map_%d", x.ID)] = reservationMap
		data[fmt.Sprintf("block_map_%d", x.ID)] = blockMap

		m.App.Session.Put(r.Context(), fmt.Sprintf("block_map_%d", x.ID), blockMap)
	}

	render.Template(w, r, "admin-calendar.page.html", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
		IntMap: intMap,
	})
}

// AdminPostReservationsCalendar handles post changes of reservation calendar
func (m *Repository) AdminPostReservationsCalendar(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	year, _ := strconv.Atoi(r.Form.Get("y"))
	month, _ := strconv.Atoi(r.Form.Get("m"))

	// process blocks
	rooms, err := m.DB.AllRooms()
	if err != nil {
		helpers.ServerError(w, err)
		return 
	}

	form := forms.New(r.PostForm)

	for _, x := range rooms {
		// Get the block map from the session. Lop through entire map and if there is 
		// an entry in the map that does not exist in the posted data, and if the 
		// restriction id > 0, then remove the block
		curMap := m.App.Session.Get(r.Context(), fmt.Sprintf("block_map_%d", x.ID)).(map[string]int)
		for name, value := range curMap {
			// 'ok' will be false if the value is not in the map
			if val, ok := curMap[name]; ok {
				// only pay attention to values > 0, and that are not in the form post
				// the rest are just placeholders for days without blocks
				if val > 0 {
					if !form.Has(fmt.Sprintf("remove_block_%d_%s", x.ID, name)) {
						err := m.DB.DeleteBlockByID(value)
						if err != nil {
							log.Println(err)
						}
					}
				}

			}
		}
	}

	// now handle new blocks: loop through post form and look for any entries with the name "add_block_"
	for name, _ := range r.PostForm {
		if strings.HasPrefix(name, "add_block") {
			exploded := strings.Split(name, "_")
			roomID, _ := strconv.Atoi(exploded[2])
			startDate, _ := time.Parse("2006-01-2", exploded[3])
			// insert a new block
			err := m.DB.InsertBlockForRoom(roomID, startDate)
			if err != nil {
				log.Println(err)
			}
		}
	}

	m.App.Session.Put(r.Context(), "flash", "Changes saved")
	http.Redirect(w, r, fmt.Sprintf("/admin/reservations-calendar?y=%d&m=%d", year, month), http.StatusSeeOther)
}

// ---------------------------------------------------------------------------------------------

// Smallchange renders the Metamask info page
func (m *Repository) PostSmallchange(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	account := r.Form.Get("eth_account")

	data := make(map[string]interface{})
	data["eth-account"] = account
	log.Println("account: ", account)

	BaseCampaignDeploy(account)

	render.Template(w, r, "metamask-info.page.html", &models.TemplateData{
		Data: data,
	})
}

func BaseCampaignDeploy(_account string) {
	//call SetUpParams() in smart.go
	client, _, nonce := ether.SetUpParams("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5", "94851db7c06aa833f3d34217e0eac7f467f5010f9a79bb9d193e14d9b23688f7")

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/9c8598071c5f4e9483ec92a18cbccbb5")
	if err != nil {
		log.Fatal(err)
	}

	//chainID, err := client.ChainID(context.Background())
	//if err != nil {
		//panic(err)
	//}

	//auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//if err != nil {
		//log.Fatal(err)
	//}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)
	//auth.GasLimit = uint64(3000000)
	//auth.GasPrice = gasPrice
	//auth.Nonce = big.NewInt(int64(nonce))

	initValue := big.NewInt(10)
	//publicKey := privateKey.Public()
	//pubKey := crypto.CreateAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9", nonce)
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
		//log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fromAddress := common.HexToAddress("0x2355B93c4551a235315e14c1e76AcDDe0Bfcc4C9")

	auth2 := bind.TransactOpts {
		From: fromAddress,
		Nonce : big.NewInt(int64(nonce)),
		Signer: nil,
		Value: big.NewInt(0),
		GasPrice: gasPrice,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit: uint64(3000000),
	}

	address, _, _, err := adf2.DeployBaseCampaign(&auth2, client, initValue)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\ncontract address: ", address.Hex())
}

/* type TransactOpts struct {
	From   common.Address // Ethereum account to send the transaction from
	Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
	Signer SignerFn       // Method to use for signing the transaction (mandatory)

Value     *big.Int // Funds to transfer along the transaction (nil = 0 = no funds)
GasPrice  *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasFeeCap *big.Int // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle)
	GasTipCap *big.Int // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle)
GasLimit  uint64   // Gas limit to set for the transaction execution (0 = estimate)

	Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)

	NoSend bool // Do all transact steps but do not send the transaction
} */