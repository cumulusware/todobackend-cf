package todos

import "github.com/gorilla/mux"

// AddRoutes adds subroutes using the URI to the given router.
func AddRoutes(r *mux.Router, uri string) {
	s := r.PathPrefix(uri).Subrouter()
	s.HandleFunc("/", DescribeAll()).Methods("OPTIONS")
	s.HandleFunc("/{key}", Describe()).Methods("OPTIONS")
	s.HandleFunc("/", ReadAll()).Methods("GET")
	s.HandleFunc("/", Create()).Methods("POST")
	s.HandleFunc("/", DeleteAll()).Methods("DELETE")
}
