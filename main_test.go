package main

import (
	"testing"
)

func TestGetRouter(t *testing.T) {
	router := getRouter()
	if router == nil {
		t.Errorf("router should not be nil")
	}
}
