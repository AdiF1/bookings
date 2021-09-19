package helpers

import (
	"math/rand"
	"time"
)

// packages
	type SomeType struct {
		TypeName string
	}
// channels
	func RandomNumber(n int) int {
		rand.Seed(time.Now().Unix())
		value := rand.Intn(n)
		return value
	}