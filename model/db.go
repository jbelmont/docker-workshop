package model

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// DBName is name of the database and CName is the name of document collection
const (
	DBName = "docker"
	CName  = "users"
)

var session *mgo.Session

func createDBSession() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
}

func getSession() *mgo.Session {
	if session == nil {
		createDBSession()
	}
	return session
}

// InitDB is to be used in main.go to initialize database
func InitDB() *mgo.Collection {
	createDBSession()
	c := NewContext()
	return c.DBCollection()
}
