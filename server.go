package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ServeHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!") // send data to client side
}

func ServeUser(state *HttpState) {
	user := User{
		Name:    "Fred",
		Created: time.Date(2007, 5, 12, 0, 0, 0, 0, time.UTC),
		Private: UserPrivate{
			EmployeeId: 12345,
		},
	}
	buf, _ := json.Marshal(user)
	state.SendBytes(buf)
}

func SetupMux() *http.ServeMux {
	var mux = http.NewServeMux()
	mux.HandleFunc("/test", ServeHello)
	mux.HandleFunc("/api/user/1", Pipe(ServeUser))
	return mux
}
