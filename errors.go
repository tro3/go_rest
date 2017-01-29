package go_rest

import "net/http"

var Errors = map[int]string{
	http.StatusForbidden: "Forbidden",
	http.StatusNotFound:  "Not Found",
}
