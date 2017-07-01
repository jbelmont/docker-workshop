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

func getDial() string {
	var dial string
	if os.Getenv("MONGO_URL") == "db" {
		dial = os.Getenv("MONGO_URL")
	} else {
		dial = "localhost"
	}
	return dial
}

// CreateDBSession creates mgo session
func CreateDBSession() *mgo.Session {
	var err error
	var dial string
	if os.Getenv("MONGO_URL") == "db" {
		dial = os.Getenv("MONGO_URL")
	} else {
		dial = "localhost"
	}
	session, err = mgo.Dial(dial)
	defer session.Close()
	anotherSession := session.Copy()
	if err != nil {
		log.Fatal("Cannot Dial Mongo: ", err)
	}
	session.SetMode(mgo.Monotonic, true)
	collection := InitDB()
	CreateInitDocument(collection)
	return anotherSession
}

// GetSession will create a session if doesn't exist else will return existing session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(getDial())
		if err != nil {
			log.Fatal("Cannot Dial Mongo: ", err)
		}
		names, _ := session.DatabaseNames()
		var dockerExists bool
		for _, name := range names {
			if name == "docker" {
				dockerExists = true
			}
		}
		if !dockerExists {
			CreateDBSession()
		}
	}
	return session
}

// GetContext get mgo collection context
func GetContext() *Context {
	session, err := mgo.Dial(getDial())
	if err != nil {
		log.Fatal("Cannot Dial Mongo: ", err)
	}
	return &Context{
		Session: session,
	}
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

// Query conducts search on mongodb collection
func Query(m *mgo.Collection, query interface{}) interface{} {
	var model Model
	return m.Find(query).One(&model)
}
