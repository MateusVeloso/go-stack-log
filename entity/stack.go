package entity

import "time"

//Stack data
type Stack struct {
	ID        ID
	Title     string
	Logs      []ID
	CreatedAt time.Time
}

//NewStack create a new Stack
func NewStack(StackTitle string) (*Stack, error) {
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

func (s *Stack) AddLog(id ID) error {
	if _, err := s.GetLog(id); err == nil {
		return ErrLogAlreadyBorrowed
	}
	s.Logs = append(s.Logs, id)
	return nil
}

//RemoveLog remove a log
func (s *Stack) RemoveLog(id ID) error {
	for i, j := range s.Logs {
		if j == id {
			s.Logs = append(s.Logs[:i], s.Logs[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

//GetLog get a log
func (s *Stack) GetLog(id ID) (ID, error) {
	for _, v := range s.Logs {
		if v == id {
			return id, nil
		}
	}
	return id, ErrNotFound
}

//Validate validate data
func (s *Stack) Validate() error {
	if s.Title == "" {
		return ErrInvalidEntity
	}
	return nil
}
