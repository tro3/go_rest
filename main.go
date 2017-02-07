package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello"))
}

func printRoute(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	t, _ := route.GetPathTemplate()
	fmt.Println(route.GetName(), t)
	return nil
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	APIRouter(router.PathPrefix("/api").Subrouter())
	//router.Walk(printRoute)

	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
