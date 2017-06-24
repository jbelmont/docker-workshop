package model

import (
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
		panic(err)
	}
}

func GetSession() *mgo.Session {
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

// CreateInitDocument initializes collection with 50 users
func CreateInitDocument(m *mgo.Collection) {
	models := GetModels()
	for _, model := range models {
		m.Insert(model)
	}
	session := GetSession()
	session.Close()
}
