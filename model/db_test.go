package model

import (
	"testing"

	"gopkg.in/mgo.v2"
)

var db *mgo.Collection

func TestInitDB(t *testing.T) {
	db = InitDB()
}

func TestCreateInitDocument(t *testing.T) {
	CreateInitDocument(db)
}
