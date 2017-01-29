package go_rest

import (
	"bytes"
	"io"
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

func compareBytes(t *testing.T, val, exp []byte, label string) {
	if !bytes.Equal(val, exp) {
		t.Errorf("Failure %s: Expected '%s', got '%s'", label, exp, val)
	}
}

func compareInt(t *testing.T, val, exp int, label string) {
	if val != exp {
		t.Errorf("Failure %s: Expected '%d', got '%d'", label, exp, val)
	}
}

func checkNil(t *testing.T, err error, label string) bool {
	if err != nil {
		t.Errorf("Failure %s: Expected nil error, got '%s'", label, err.Error())
		return true
	}
	return false
}

func genRequest(method, URL string, body io.Reader) (res *httptest.ResponseRecorder, req *http.Request, err error) {
	res = httptest.NewRecorder()
	req, err = http.NewRequest(method, URL, body)
	return
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
