package rest

import "testing"

func TestReferenceMarshall(t *testing.T) {
	ref := Reference{
		Collection: "test",
		Id:         4,
	}
	buf, err := ref.MarshalJSON()
	if checkNil(t, err, "REF0") {
		return
	}
	compareString(t, string(buf), "Test", "REF1")
}

func TestReferenceUnmarshall(t *testing.T) {
	var ref = new(Reference)
	ref.UnmarshalJSON([]byte("Test"))
	compareString(t, ref.Collection, "test", "REF2")
	compareInt(t, ref.Id, 4, "REF3")
}
