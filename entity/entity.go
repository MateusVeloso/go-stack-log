package entity

import "github.com/google/uuid"

//ID entity
type ID = uuid.UUID

//NewID create new entity ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
