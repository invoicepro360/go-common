package cutils

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/invoicepro360/go-common/ctemplates"
)

// FailedResponse provides response for failed requests incase of errors
func FailedResponse(r *http.Request, w http.ResponseWriter, httpStatus int, message string, errorMessage interface{}) {

	var badResponse ctemplates.BadResponse
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

	var goodResponse ctemplates.GoodResponse
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

	var goodResponse ctemplates.GoodResponseWithPagination
	goodResponse.Status = httpStatus
	goodResponse.Meta.TotalResults = totalResults
	if totalResults > 0 {
		TotalPages := math.Ceil(float64(totalResults / size))
		goodResponse.Meta.TotalPages = int(TotalPages)
	} else {
		goodResponse.Meta.TotalPages = 0
	}
	goodResponse.Meta.Page = page
	goodResponse.Meta.PageSize = size
	goodResponse.Data = data

	// write response
	json.NewEncoder(w).Encode(goodResponse)

}
