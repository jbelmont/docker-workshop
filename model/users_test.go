package model

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	users := GetUsers()
	if len(users) != 50 {
		t.Errorf("Should have returned a list of 50 users")
	}
}

func TestGetUsersById(t *testing.T) {
	user, err := GetUserByID(3)
	if err != nil {
		t.Fatal(err)
	}
	if user.FirstName != "Jonathan" {
		t.Errorf("should have returned the first name of %s", user.FirstName)
	}
}
