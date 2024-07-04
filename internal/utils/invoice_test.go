package utils

import (
	"regexp"
	"testing"
)

func TestGenerateInvoice(t *testing.T) {
	invoiceNumber := GenerateInvoice()
	matched, err := regexp.MatchString(`^INV/\d{8}/MPL/\d{8}$`, invoiceNumber)
	if err != nil {
		t.Fatalf("Error matching regex: %v", err)
	}
	if !matched {
		t.Errorf("Invoice number %s does not match the expected format", invoiceNumber)
	}
}
