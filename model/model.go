package model

// Create method for the store that creates another Model
func (m *Model) Create() Model {
	return *m
}

// GetAll method returns all the Models
func (m *Model) GetAll() ([]Model, error) {
	return []Model{}, nil
}

// GetByID method returns Model with given id
func (m *Model) GetByID() (Model, error) {
	panic("implement me")
}

// Update method updates the given Model
func (m *Model) Update() (Model, error) {
	panic("implement me")
}

// Delete method removes the given Model
func (m *Model) Delete() (Model, error) {
	panic("implement me")
}

// Models type which is slice of Model
type Models []Model

// New returns new Model
func New() *Model {
	return &Model{}
}
