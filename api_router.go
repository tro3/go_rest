package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var endpoints = []Endpoint{
	UserEndpoint,
}

type apiCommand struct {
	name    string
	method  string
	pattern string
	fn      http.Handler
}

func apiCommands(e Endpoint) []apiCommand {
	return []apiCommand{
		{"List", "GET", "/", e.GetList},
		{"Item", "GET", "/{id}", e.GetItem},
		{"Create", "POST", "/", e.CreateItem},
		{"Edit", "PUT", "/{id}", e.EditItem},
		{"Delete", "DELETE", "/{id}", e.DeleteItem},
	}
}

func APIRouter(router *mux.Router) {
	for _, endpoint := range endpoints {
		subrouter := router.PathPrefix(endpoint.Path).Subrouter()
		for _, cmd := range apiCommands(endpoint) {
			name := endpoint.Name + cmd.name
			fn := cmd.fn
			// Add middleware here
			fn = Logger(fn, name)

			subrouter.
				Methods(cmd.method).
				Path(cmd.pattern).
				Name(name).
				Handler(fn)
		}
	}
}
