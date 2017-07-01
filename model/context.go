package model

import (
	mgo "gopkg.in/mgo.v2"
)

// Context is an mgo session
type Context struct {
	Session *mgo.Session
}

// Close the context session
func (c *Context) Close() {
	c.Session.Close()
}

// DBCollection returns mgo collection
func (c *Context) DBCollection() *mgo.Collection {
	return c.Session.DB(DBName).C(CName)
}

// MongoSetup returns context
func MongoSetup() *Context {
	session := GetSession()
	c := &Context{
		Session: session,
	}
	return c
}
