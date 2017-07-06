package redis

import (
	"strconv"
	"testing"

	"github.com/jbelmont/docker-workshop/model"
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

func TestSetHash(t *testing.T) {
	testUser := []model.UserModel{
		model.UserModel{
			ID: 1, FirstName: "Timothy", LastName: "Cox", Email: "tcox0@dion.ne.jp", Gender: "Male",
		},
		model.UserModel{
			ID: 2, FirstName: "Sean", LastName: "Medina", Email: "smedina1@addthis.com", Gender: "Male",
		},
	}
	for key, user := range testUser {
		setName := "testuser:" + strconv.Itoa(key)
		SetHash(setName, user)
	}
}

func TestGetHash(t *testing.T) {
	reply, err := GetHashAll("user:1")
	if err != nil {
		t.Error("should be nil")
	}
	_, err2 := reply.Map()
	if err2 != nil {
		t.Error("error converting")
	}
}

func TestGetKeysWithPattern(t *testing.T) {
	reply, _ := GetKeys("user:*")
	users, _ := reply.Array()
	for _, k := range users {
		key, _ := k.String()
		rep, _ := GetHashAll(key)
		_, err := rep.Map()
		if err != nil {
			t.Error("something went wrong")
		}
	}
}
