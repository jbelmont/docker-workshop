package model

import mgo "gopkg.in/mgo.v2"

type Context struct {
	Session *mgo.Session
}

func (c *Context) Close() {
	c.Session.Close()
}

func (c *Context) DBCollection() *mgo.Collection {
	return c.Session.DB(DBName).C(CName)
}

func NewContext() *Context {
	session := getSession()
	c := &Context{
		Session: session,
	}
	return c
}
