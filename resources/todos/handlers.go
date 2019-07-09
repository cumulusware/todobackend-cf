package todos

import (
	"encoding/json"
	"log"
	"net/http"
	"path"

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
func ReadAll(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseURL := createURL(r)
		todos, err := ds.GetAll(baseURL)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, todos)
	}
}

// Read handles the GET method to list all todos.
func Read(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)
		url := createURL(r)
		todo, err := ds.GetByID(id, url)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, todo)
	}
}

// Create handles the POST method to create a new todo.
func Create(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&t); err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		id, err := ds.Create(&t)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		baseURL := createURL(r)
		t.URL = baseURL + id
		w.Header().Set("Location", baseURL+id)

		helpers.RespondWithJSON(w, http.StatusCreated, t)
	}
}

// DeleteAll handles the DELETE method to delete all todos.
func DeleteAll(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := ds.DeleteAll(); err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusNoContent, "")
	}
}

// Delete handles the DELETE method to delete a todo.
func Delete(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := path.Base(r.URL.Path)
		if err := ds.DeleteByID(id); err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		helpers.RespondWithJSON(w, http.StatusNoContent, "")
	}
}

// Update handles the PATCH method to update a portion of a todo.
func Update(ds DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get current todo using ID.
		id := path.Base(r.URL.Path)
		url := createURL(r)
		todo, err := ds.GetByID(id, url)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Update the todo.
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		err = ds.UpdateByID(id, &todo)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		helpers.RespondWithJSON(w, http.StatusCreated, todo)
	}
}

func createURL(r *http.Request) string {
	protocol, err := helpers.Protocol(r.Host)
	if err != nil {
		protocol = "https://"
		log.Printf("Error determining protocol for host %s. Defaulting to https://", r.Host)
	}
	return protocol + r.Host + r.URL.String()
}
