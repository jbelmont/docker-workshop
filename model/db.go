package model

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// DBName is name of the database and CName is the name of document collection
const (
	DBName = "docker"
	CName  = "users"
)

var session *mgo.Session

func CreateDBSession() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:   []string{os.Getenv("MONGO_URL")},
		Timeout: 60 * time.Second,
	}
	db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal("Cannot Dial Mongo: ", err)
	}
	defer db.Close()
	db.SetMode(mgo.Monotonic, true)
}

func GetSession() *mgo.Session {
	if session == nil {
		CreateDBSession()
	}
	return session
}

// InitDB is to be used in main.go to initialize database
func InitDB() *mgo.Collection {
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
	defer session.Close()
}
