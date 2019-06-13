package helpers

import "net/http"

// RespondWithOptions responds to an HTTP request with allowed options.
func RespondWithOptions(w http.ResponseWriter, options string) {
	w.Header().Set("Allow", options)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
