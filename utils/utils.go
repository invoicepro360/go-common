package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/invoicepro360/go-common/config"
	"github.com/invoicepro360/go-common/templates"
)

// FailedResponse provides response for failed requests incase of errors
func FailedResponse(r *http.Request, w http.ResponseWriter, httpStatus int, message string, errorMessage interface{}) {

	var badResponse templates.BadResponse
	badResponse.Status = httpStatus
	badResponse.Message = message

	switch t := errorMessage.(type) {
	case string:

		badResponse.Error = t
		badResponse.ValidationErrors = nil

	default:
		badResponse.Error = ""
		badResponse.ValidationErrors = errorMessage
	}

	// set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	// write response
	json.NewEncoder(w).Encode(badResponse)

}

// SuccessResponse provides response for successful requests
func SuccessResponse(r *http.Request, w http.ResponseWriter, httpStatus int, message string, data interface{}) {

	var goodResponse templates.GoodResponse
	goodResponse.Status = httpStatus
	goodResponse.Message = message
	goodResponse.Data = data

	// set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	// write response
	json.NewEncoder(w).Encode(goodResponse)

}

// SuccessResponseResults provides response for successful requests
func SuccessResponseResults(r *http.Request, w http.ResponseWriter, httpStatus int, totalResults int, page int, size int, data interface{}) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	var goodResponse templates.GoodResponse
	goodResponse.Meta.TotalResults = totalResults
	if totalResults > 0 {
		goodResponse.Meta.TotalPages = int(totalResults / size)
	} else {
		goodResponse.Meta.TotalPages = 0
	}
	goodResponse.Meta.Page = page
	goodResponse.Meta.PageSize = size
	goodResponse.Data = data

	// write response
	json.NewEncoder(w).Encode(goodResponse)
	return
}

// NoRouteFoundHandler handles cases where an undefined routes are requested
func NoRouteFoundHandler(w http.ResponseWriter, r *http.Request) {
	errorMessage := fmt.Sprintf("Invalid endpoint request (%v)", r.URL.Path)
	FailedResponse(r, w, http.StatusNotFound, "", errorMessage)
}

// HealthCheckHandler handles /healthcheck route
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var items []templates.HealthCheckItem
	var item templates.HealthCheckItem
	var hasError = false

	// test container
	item.Name = "CONTAINER"
	item.IsHealthy = true
	item.Message = "Container is connected"
	items = append(items, item)

	// test database connection
	conn, err := net.DialTimeout("tcp", config.DBHost+":"+config.DBPort, 2*time.Second)
	if err != nil {
		item.Name = "DATABASE"
		item.IsHealthy = false
		item.Message = "Database connection failed"
		items = append(items, item)
		hasError = true
	} else {
		item.Name = "DATABASE"
		item.IsHealthy = true
		item.Message = "Database is connected"
		items = append(items, item)
		defer conn.Close()
	}

	// set headers
	w.Header().Set("Content-Type", "application/json")
	if hasError == false {
		// success (200)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
	} else {
		// failed (503)
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(items)
	}
}
