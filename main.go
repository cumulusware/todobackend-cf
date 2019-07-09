package main

import (
	"context"
	"log"
	"net/http"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/cumulusware/todobackend-cf/data/couchdb"
	"github.com/flimzy/kivik"
	_ "github.com/go-kivik/couchdb"
	"github.com/joho/godotenv"
)

const (
	dbName = "tododb"
)

func main() {

	// Get Cloudant credentials from Cloud Foundry or locally.
	var cloudantURL string
	if appEnv, err := cfenv.Current(); err != nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		cloudantURL = os.Getenv("CLOUDANT_URL")
	} else {
		cloudantService, err := appEnv.Services.WithName(dbName)
		if err != nil {
			log.Fatal("Could not find cloudant db.")
		}
		// The Cloudant URL will contain the username and password if it was
		// provisioned using both the Legacy and IAM credentials.
		cloudantURL = cloudantService.Credentials["url"].(string)
	}

	// Create the data store.
	client, err := kivik.New(context.TODO(), "couch", cloudantURL)
	if err != nil {
		log.Fatalf("Failed connecting to Cloudant database: %s", err)
	}
	ds, err := couchdb.NewDataStore(context.TODO(), client)
	if err != nil {
		log.Fatalf("error creating datastore: %s", err)
	}

	// Get port or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create routes with injected data store and CORS. Then start server.
	r := createRoutes(ds)
	c := setupCors()
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
