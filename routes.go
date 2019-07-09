package main

import (
	"github.com/cumulusware/todobackend-cf/data/couchdb"
	"github.com/cumulusware/todobackend-cf/resources/todos"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createRoutes(ds *couchdb.Store) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	todos.AddRoutes(r, ds.Todos, "/api/todos")
	return r
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "OPTIONS", "POST", "DELETE", "PUT", "PATCH"},
		AllowedHeaders:   []string{"accept", "content-type"},
		Debug:            true,
	})
}
