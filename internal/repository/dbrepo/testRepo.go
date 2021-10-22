package dbrepo

import (
	"time"

	"github.com/AdiF1/solidity/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {

	return true
}

// InsertReservation inserts a reservation into the DB
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil

}

// InsertRoomRestriction inserts a room restriction into the db
func(m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	
	return nil
}
// SearchAvailabilityByDatesAndRoomID returns true if availability exists for roomID, otherwise false
func (m *testDBRepo) SearchAvailabilityByDatesAndRoomID(start, end time.Time, roomID int) (bool, error) {
	
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of rooms, if any are available for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}
// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	
	var room models.Room

	return room, nil

}