package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// Get port or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create routes with CORS. Then start server.
	r := createRoutes()
	c := setupCors()
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
