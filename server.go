package go_rest

import (
	"encoding/json"
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

func SetupRouter() *Router {
	var router = NewRouter()
	router.Get("/test", Pipe(ServeHello))
	router.Get("/api/user/1", Pipe(ServeUser))
	return router
}
