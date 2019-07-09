package todos

import "github.com/gorilla/mux"

// AddRoutes adds subroutes using the URI to the given router.
func AddRoutes(r *mux.Router, ds DataStore, uri string) {
	s := r.PathPrefix(uri).Subrouter()
	s.HandleFunc("/", DescribeAll()).Methods("OPTIONS")
	s.HandleFunc("/", ReadAll(ds)).Methods("GET")
	s.HandleFunc("/", Create(ds)).Methods("POST")
	s.HandleFunc("/", DeleteAll(ds)).Methods("DELETE")
	s.HandleFunc("/{key}", Describe()).Methods("OPTIONS")
	s.HandleFunc("/{key}", Read(ds)).Methods("GET")
	s.HandleFunc("/{key}", Update(ds)).Methods("PATCH")
	s.HandleFunc("/{key}", Delete(ds)).Methods("DELETE")
}
