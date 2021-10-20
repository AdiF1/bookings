package main

import (
	"fmt"
	"testing"

	"github.com/AdiF1/solidity/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	testMux := routes(&app)

	switch v := testMux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))


	}	


}