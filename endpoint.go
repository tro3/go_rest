package main

import "net/http"

type Endpoint struct {
	Name       string
	Path       string
	GetList    http.HandlerFunc
	GetItem    http.HandlerFunc
	CreateItem http.HandlerFunc
	EditItem   http.HandlerFunc
	DeleteItem http.HandlerFunc
}
