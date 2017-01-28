package rest

import (
	"fmt"
	"net/http"
)

var Errors = map[int]string{
	http.StatusForbidden: "Forbidden",
	http.StatusNotFound:  "Not Found",
}

type HttpState struct {
	Req *http.Request
	Res http.ResponseWriter
	Err int
}

type HttpStateFunc func(*HttpState)

func Pipe(fns ...HttpStateFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		state := &HttpState{req, res, 0}
		for _, fn := range fns {
			fn(state)
			if state.Err != 0 {
				res.Write([]byte(fmt.Sprintf("%d %s", state.Err, Errors[state.Err])))
				return
			}
		}
	}
}
