package go_rest

import (
	"fmt"
	"strings"
	"testing"
)

func TestRouterBasics(t *testing.T) {
	fn := Pipe(func(state *HttpState) {
		state.SendString(state.Req.Method)
	})

	router := NewRouter()
	router.Get("/", fn)
	router.Post("/", fn)
	router.Put("/", fn)
	router.Delete("/", fn)

	verbs := []string{"GET", "POST", "PUT", "DELETE"}

	for _, verb := range verbs {
		res, req, err := genRequest(verb, "/", strings.NewReader(""))
		checkNil(t, err, "ROUTE0")
		router.ServeHTTP(res, req)
		compareString(t, res.Body.String(), verb, fmt.Sprintf("ROUTE1/%s", verb))
	}
}

func TestRouterNotFound(t *testing.T) {
	fn := Pipe(func(state *HttpState) {
		state.SendString(state.Req.Method)
	})

	router := NewRouter()
	router.Get("/", fn)
	router.Post("/", fn)
	router.Put("/", fn)
	router.Delete("/", fn)

	verbs := []string{"GET", "POST", "PUT", "DELETE"}

	for _, verb := range verbs {
		res, req, err := genRequest(verb, "/a", strings.NewReader(""))
		checkNil(t, err, "ROUTE2")
		router.ServeHTTP(res, req)
		compareString(t, res.Body.String(), "404 Not Found", fmt.Sprintf("ROUTE3/%s", verb))
	}
}
