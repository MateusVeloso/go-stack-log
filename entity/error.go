package entity

import "errors"

//ErrNotValidTypeLog not valid type log
var ErrNotValidTypeLog = errors.New("Invalid Type Log")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrLogAlreadyBorrowed cannot borrowed
var ErrLogAlreadyBorrowed = errors.New("Logs already borrowed")

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")
