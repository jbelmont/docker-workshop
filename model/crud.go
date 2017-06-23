package model

import "gopkg.in/mgo.v2/bson"

// Model is a mgo struct data type
type Model struct {
	ID        bson.ObjectId `bson:"_id, omitempty"`
	FirstName string
	LastName  string
	Email     string
	Gender    string
}

// Store has CRUD methods for mongodb
type Store interface {
	Create(m *Model) Model
	GetAll() ([]Model, error)
	GetByID(m *Model) (Model, error)
	Update(m *Model) (Model, error)
	Delete(m *Model) (Model, error)
}
