package rest

import "fmt"

type Reference struct {
	Collection string
	Id         int
}

func (self Reference) MarshalJSON() (json []byte, err error) {
	fmt.Println(self)
	err = nil
	return []byte("Test"), err
}

func (self *Reference) UnmarshalJSON(json []byte) error {
	self.Collection = "test"
	self.Id = 4
	return nil
}
