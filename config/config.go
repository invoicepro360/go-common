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
		"MM/DD/YY":    "1/2/2006",
		"DD/MM/YY":    "1/2/2006",
		"YY/MM/DD":    "2006/2/1",
		"Month D, YY": "January 2, 2006",
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
