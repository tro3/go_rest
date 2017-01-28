package rest

import (
	"fmt"
	"net/http"
)

func ServeHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!") // send data to client side
}

func SetupMux() *http.ServeMux {
	var mux = http.NewServeMux()
	mux.HandleFunc("/test", ServeHello)
	return mux
}
