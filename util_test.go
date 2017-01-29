package go_rest

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
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
