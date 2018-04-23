package gohostid

import "testing"

func TestSimpleCall(t *testing.T) {
	id, err := GetID()
	if err != nil {
		t.Fail()
	}
	t.Log(id)
}
