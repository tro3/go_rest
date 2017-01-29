package go_rest

import (
	"encoding/json"
	"net/http"
	"time"
)

func ServeHello(state *HttpState) {
	state.SendString("Hello!")
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
	state.SendJsonBytes(buf)
}

func SetupMux() *http.ServeMux {
	var mux = http.NewServeMux()
	mux.HandleFunc("/test", Pipe(ServeHello))
	mux.HandleFunc("/api/user/1", Pipe(ServeUser))
	return mux
}
