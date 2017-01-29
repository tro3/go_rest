package go_rest

import (
	"strings"
	"testing"
)

func TestSendString(t *testing.T) {
	f1 := func(state *HttpState) {
		state.SendString(state.Req.URL.String())
	}

	res, req, err := genRequest("GET", "/fred", strings.NewReader(""))
	checkNil(t, err, "STATE0")

	Pipe(f1)(res, req)
	compareString(t, res.Body.String(), "/fred", "STATE1")
}

func TestSendBytes(t *testing.T) {
	f1 := func(state *HttpState) {
		state.SendBytes([]byte(state.Req.URL.String()))
	}

	res, req, err := genRequest("GET", "/george", strings.NewReader(""))
	checkNil(t, err, "STATE1")

	Pipe(f1)(res, req)
	compareBytes(t, res.Body.Bytes(), []byte("/george"), "STATE2")
}

func TestSendJsonBytes(t *testing.T) {
	f1 := func(state *HttpState) {
		state.SendJsonBytes([]byte(`{"name":"Bob"}`))
	}

	res, req, err := genRequest("GET", "/bob", strings.NewReader(""))
	checkNil(t, err, "STATE3")

	Pipe(f1)(res, req)
	compareString(t, res.Body.String(), `{"name":"Bob"}`, "STATE4")
	compareString(t, res.Header().Get("Content-Type"), "application/json", "STATE5")
}
