package config

import (
	"os"
)

var (
	// IsDebug provide flags for logging
	IsDebug bool = false

	// DBUser sets db connection username
	DBUser string
	// DBPassword sets db connection password
	DBPassword string
	// DBHost sets db connection host
	DBHost string
	// DBPort sets db connection port
	DBPort string
	// DBName sets db connection database name
	DBName string
	// SecuritySalt sets security salt used to encode/decode user password
	SecuritySalt string
	// UserJwt set user information returned by jwt token validation

	PdfInvoiceTemplates string

	DateFormat = map[string]string{
		"MM/DD/YYYY":   "01/02/2006",
		"DD/MM/YYYY":   "02/01/2006",
		"YYYY/MM/DD":   "2006/01/02",
		"MMMM D, YYYY": "January 2, 2006",
	}
)

// Initialize set config defaults
func Initialize() {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	SecuritySalt = os.Getenv("SECURITY_SALT")
	PdfInvoiceTemplates = os.Getenv("PDF_INVOICE_TEMPLATES")
}
