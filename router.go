package go_rest

import (
	"fmt"
	"net/http"
)

type routeMap map[string]http.HandlerFunc

type Router struct {
	getPaths    routeMap
	postPaths   routeMap
	putPaths    routeMap
	deletePaths routeMap
}

func NewRouter() *Router {
	var router = Router{
		getPaths:    make(routeMap),
		postPaths:   make(routeMap),
		putPaths:    make(routeMap),
		deletePaths: make(routeMap),
	}

	return &router
}

func (self *Router) Get(path string, fn http.HandlerFunc) {
	self.getPaths[path] = fn
}

func (self *Router) Post(path string, fn http.HandlerFunc) {
	self.postPaths[path] = fn
}

func (self *Router) Put(path string, fn http.HandlerFunc) {
	self.putPaths[path] = fn
}

func (self *Router) Delete(path string, fn http.HandlerFunc) {
	self.deletePaths[path] = fn
}

func (self *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleMethod(self.getPaths, w, r)
	case "POST":
		handleMethod(self.postPaths, w, r)
	case "PUT":
		handleMethod(self.putPaths, w, r)
	case "DELETE":
		handleMethod(self.deletePaths, w, r)
	default:
		w.Write([]byte("Unsupported command"))
	}
}

func handleMethod(routes routeMap, w http.ResponseWriter, r *http.Request) {
	fn, found := routes[r.URL.String()]
	if found {
		fn(w, r)
	} else {
		w.Write([]byte(fmt.Sprintf("%d %s", 404, Errors[404])))
	}
}
