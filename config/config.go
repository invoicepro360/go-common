package config

import "os"

const (
	// BaseURL sets api in url path
	BaseURL string = "/api"
	// APIVersion set api version in path
	APIVersion string = "/v1"
	// StaticFilesDir defines the dir for static files
	StaticFilesDir string = "docs"
	// DefaultPageSize set number of rows to return
	DefaultPageSize int = 15
	// DefaultPage set which page to return
	DefaultPage int = 1
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
)

// Initialize set config defaults
func Initialize() {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	SecuritySalt = os.Getenv("SECURITY_SALT")
}
