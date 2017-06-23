package model

import (
	"testing"
)

var model Model

func init() {
	model = Model{
		ID:        "1",
		FirstName: "John",
		LastName:  "Rambo",
		Email:     "john.rambo@badass.net",
		Gender:    "Male",
	}
}

func TestCreate(t *testing.T) {
	actual := model.Create()
	if actual.Email != "john.rambo@badass.net" {
		t.Errorf("%s should be ", actual.Email)
	}
}
