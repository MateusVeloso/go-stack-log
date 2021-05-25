package entity

import "time"

var validTypes = map[int]string{
	1: "INFO",
	2: "LOG",
	3: "WARN",
	4: "ERROR",
}

//Stack data
type Stack struct {
	ID        ID
	Title     string
	Logs      []Log
	CreatedAt time.Time
}

//Log data
type Log struct {
	ID        ID
	Type      string
	Text      string
	CreatedAt time.Time
}

func NewStackLog(StackTitle string) (*Stack, error) {
	stack := &Stack{
		ID:        NewID(),
		Title:     StackTitle,
		CreatedAt: time.Now(),
	}
	if err := stack.Validate(); err != nil {
		return nil, err
	}
	return stack, nil
}

func (s *Stack) Validate() error {
	if s.Title == "" {
		return ErrInvalidEntity
	}
	return nil
}

func (log *Log) SetType(t int) error {
	if tp, ok := validTypes[t]; ok {
		log.Type = tp
		return nil
	}
	return ErrNotValidTypeLog
}
