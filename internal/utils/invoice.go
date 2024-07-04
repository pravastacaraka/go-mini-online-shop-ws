package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoice() string {
	// Get the current date in YYYYMMDD format
	currentDate := time.Now().Format("20060102")

	// Create a new random source and generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random 8-digit number
	randomNumber := r.Intn(90000000) + 10000000

	return fmt.Sprintf("INV/%s/MPL/%08d", currentDate, randomNumber)
}
