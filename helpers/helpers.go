package helpers

import (
	"fmt"
	//"math/rand"
	"net/http"
	"runtime/debug"
	//"time"

	"github.com/AdiF1/solidity/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status )
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

// packages and channels
	/* type SomeType struct {
		TypeName string
	}
	func RandomNumber(n int) int {
		rand.Seed(time.Now().Unix())
		value := rand.Intn(n)
		return value
	} */