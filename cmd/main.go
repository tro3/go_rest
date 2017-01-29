package main

import "net/http"
import "github.com/tro3/go_rest"

func main() {
	http.ListenAndServe(":8080", go_rest.SetupMux())
}
