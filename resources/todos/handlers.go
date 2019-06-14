package todos

import (
	"net/http"

	"github.com/cumulusware/todobackend-cf/helpers"
)

// DescribeAll handles the OPTIONS method for the todos/ endpoint.
func DescribeAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithOptions(w, "GET,POST,DELETE,OPTIONS")
	}
}

// Describe handles the OPTIONS method for the todos/ endpoint.
func Describe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpers.RespondWithOptions(w, "GET,PATCH,DELETE,OPTIONS")
	}
}

// ReadAll handles the GET method to list all todos.
func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "First task",
		}
		helpers.RespondWithJSON(w, http.StatusOK, todo)
	}
}

// Create handles the POST method to create a new todo.
func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "a todo",
		}
		helpers.RespondWithJSON(w, http.StatusOK, todo)
	}
}
