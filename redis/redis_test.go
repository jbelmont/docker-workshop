package redis

import (
	"testing"
)

func TestSetKey(t *testing.T) {
	SetKey("names", []string{
		"billy", "john", "will",
	})
	reply, err := GetKey("names")
	if err != nil {
		t.Errorf("something wrong happened")
	}
	expected := "[billy john will]"
	actual, _ := reply.String()
	if actual != expected {
		t.Errorf("should return %v", expected)
	}
}
