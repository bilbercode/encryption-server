package server

import "github.com/gorilla/mux"

func NewRouter(routes []Route) (*mux.Router, error) {

	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
	return router, nil
}