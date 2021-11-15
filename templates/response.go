package templates

// HealthCheckItem stores health check data
type HealthCheckItem struct {
	Name      string `json:"name"`
	IsHealthy bool   `json:"is_healthy"`
	Message   string `json:"message"`
}

// BadResponse stores invalid response data
type BadResponse struct {
	Status           int         `json:"status"`
	Message          string      `json:"message"`
	Error            string      `json:"error"`
	ValidationErrors interface{} `json:"validation_errors"`
}

// GoodResponse stores valid response data
type GoodResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
