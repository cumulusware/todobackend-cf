package todos

import "github.com/gorilla/mux"

// AddRoutes add subroutes using the URI to the given router.
func AddRoutes(r *mux.Router, uri string) {
	s := r.PathPrefix(uri).Subrouter()
	s.HandleFunc("/", DescribeAll()).Methods("OPTIONS")
	s.HandleFunc("/{key}", Describe()).Methods("OPTIONS")
}
