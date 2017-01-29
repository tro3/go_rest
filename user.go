package go_rest

import "time"

type User struct {
	Name    string
	Created time.Time
	Private UserPrivate
}

type UserPrivate struct {
	EmployeeId int
}
