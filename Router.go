package Chit_Chat_Server

import "github.com/gorilla/mux"

func NewRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.HandlerFunc)
	}
	return router
}
