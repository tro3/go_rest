package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var UserEndpoint = Endpoint{
	Name:       "User",
	Path:       "/users",
	GetList:    GetUserList,
	GetItem:    GetUser,
	CreateItem: CreateUser,
	EditItem:   EditUser,
	DeleteItem: DeleteUser,
}

var users = []User{
	{1, "Bob"},
	{2, "Fred"},
}

func GetUserList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(users)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(users[0])
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
}

func EditUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
}
