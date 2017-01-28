package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func compareString(t *testing.T, val, exp, label string) {
	if val != exp {
		t.Errorf("Failure %s: Expected '%s', got '%s'", label, exp, val)
	}
}

func TestPipe(t *testing.T) {
	f1 := func(state *HttpState) {
		state.Res.Write([]byte(state.Req.Method + " "))
	}
	f2 := func(state *HttpState) {
		state.Res.Write([]byte(state.Req.URL.String()))
	}

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/fred", strings.NewReader(""))
	if err != nil {
		t.Error(err)
		return
	}

	Pipe(f1, f2)(res, req)
	compareString(t, res.Body.String(), "GET /fred", "PIPE1")
}

func TestPipeError1(t *testing.T) {
	f1 := func(state *HttpState) {
		state.Err = http.StatusForbidden
	}
	f2 := func(state *HttpState) {
		state.Res.Write([]byte(state.Req.URL.String()))
	}

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/fred", strings.NewReader(""))
	if err != nil {
		t.Error(err)
		return
	}

	Pipe(f1, f2)(res, req)
	compareString(t, res.Body.String(), "403 Forbidden", "PIPE2")
}

func TestPipeError2(t *testing.T) {
	f1 := func(state *HttpState) {
		state.Res.Write([]byte(state.Req.Method + " "))
	}
	f2 := func(state *HttpState) {
		state.Err = http.StatusNotFound
	}

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/fred", strings.NewReader(""))
	if err != nil {
		t.Error(err)
		return
	}

	Pipe(f1, f2)(res, req)
	compareString(t, res.Body.String(), "GET 404 Not Found", "PIPE3")
}
