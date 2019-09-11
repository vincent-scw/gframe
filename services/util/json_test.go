package util

import (
	"testing"
)

type TestModel struct {
	ID int		`json:"id"`
	Name string `json:"name"`
}

func TestToJSON(t *testing.T) {
	tm := TestModel{ID: 1, Name: "test"}
	str := ToJSON(tm)
	if string(str) != `{"id":1,"name":"test"}` {
		t.Error("marshal wrong")
	}
}

func TestToModel(t *testing.T) {
	tm := &TestModel{}
	ToModel([]byte(`{"id": 1, "name":"test"}`), tm)
	if tm.ID != 1 || tm.Name != "test" {
		t.Error("unmarshal wrong")
	}
}