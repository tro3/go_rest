package go_rest

import "net/http"

type HttpState struct {
	Req *http.Request
	Res http.ResponseWriter
	Err int
}

func (self *HttpState) SendString(body string) {
	self.Res.Write([]byte(body))
}

func (self *HttpState) SendBytes(body []byte) {
	self.Res.Write(body)
}

func (self *HttpState) SendJsonBytes(body []byte) {
	self.Res.Header().Add("Content-Type", "application/json")
	self.Res.Write(body)
}

type HttpStateFunc func(*HttpState)
