package model

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// DBName is name of the database and CName is the name of document collection
const (
	DBName = "docker"
	CName  = "users"
)

var session *mgo.Session

// CreateDBSession creates mgo session
func CreateDBSession() {
	var err error
	var dial string
	if os.Getenv("MONGO_URL") == "db" {
		dial = os.Getenv("MONGO_URL")
	} else {
		dial = "localhost"
	}
	session, err = mgo.Dial(dial)
	defer session.Close()
	if err != nil {
		log.Fatal("Cannot Dial Mongo: ", err)
	}
	session.SetMode(mgo.Monotonic, true)
	collection := InitDB()
	CreateInitDocument(collection)
}

// GetSession will create a session if doesn't exist else will return existing session
func GetSession() *mgo.Session {
	if session == nil {
		CreateDBSession()
	}
	return session
}

// InitDB is to be used in main.go to initialize database
func InitDB() *mgo.Collection {
	c := MongoSetup()
	return c.DBCollection()
}

// CreateInitDocument initializes collection with 50 users
func CreateInitDocument(m *mgo.Collection) {
	models := GetModels()
	for _, model := range models {
		m.Insert(model)
	}
}
