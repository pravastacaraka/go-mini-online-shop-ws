package utils

import (
	"fmt"
	"strings"
)

// IDR converts an integer to Indonesian Rupiah format (e.g., Rp1.000.000)
func IDR(amount any) string {
	// Convert the integer to a string
	amountStr := fmt.Sprintf("%d", amount)

	// Reverse the string to facilitate inserting dots every three digits
	reversed := reverseString(amountStr)

	// Insert dots every three digits
	var parts []string
	for i := 0; i < len(reversed); i += 3 {
		end := i + 3
		if end > len(reversed) {
			end = len(reversed)
		}
		parts = append(parts, reversed[i:end])
	}

	// Join the parts with dots and reverse back
	formatted := strings.Join(parts, ".")
	formatted = reverseString(formatted)

	// Prefix with "Rp" and return
	return "Rp" + formatted
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
