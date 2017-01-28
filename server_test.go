package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(SetupMux())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/test")
	if checkNil(t, err, "BASIC0") {
		return
	}

	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if checkNil(t, err, "BASIC1") {
		return
	}

	compareString(t, string(resp), "Hello!", "BASIC2")
}

func TestUser(t *testing.T) {
	ts := httptest.NewServer(SetupMux())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/user/1")
	if checkNil(t, err, "USER0") {
		return
	}

	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if checkNil(t, err, "USER1") {
		return
	}

	exp := `{"Name":"Fred","Created":"2007-05-12T00:00:00Z","Private":{"EmployeeId":12345}}`
	compareString(t, string(resp), exp, "USER2")
}
