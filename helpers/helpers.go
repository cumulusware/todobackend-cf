package helpers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// RespondWithOptions responds to an HTTP request with allowed options.
func RespondWithOptions(w http.ResponseWriter, options string) {
	w.Header().Set("Allow", options)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// RespondWithJSON writes an HTTP response using the given HTTP status code and
// transforms the given data into JSON.
func RespondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	content, err := json.Marshal(data)
	if string(content) == "null" {
		content = []byte("[]")
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(content)
}

// RespondWithError responds to an HTTP request with an error.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// Protocol determines the protocol based on the hostname. If localhost, the
// protocol is set to "http://". Otherwise, the protocol is set to "https://".
// This is needed return a URL with the Todo that can handle CORS.
func Protocol(host string) (string, error) {
	hostOnly, _, err := net.SplitHostPort(host)
	if err != nil {
		return "", fmt.Errorf("error splitting host and post: %s", host)
	}
	protocol := "https://"
	if hostOnly == "localhost" {
		protocol = "http://"
	}
	return protocol, nil
}
