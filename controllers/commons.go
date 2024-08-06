package controllers

// ErrorResponse represents a standard error response structure
// @Description Standard error response
// @Param message query string true "Error message"
// @Success 400 {object} ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}