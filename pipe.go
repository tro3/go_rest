package go_rest

import (
	"fmt"
	"net/http"
)

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
