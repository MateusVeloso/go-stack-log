package entity

import "time"

//Log data
type Log struct {
	ID        ID
	Type      string
	Text      string
	CreatedAt time.Time
}

//Log data
func NewLog(t int, text string) (*Log, error) {
	l := &Log{
		ID:        NewID(),
		Text:      text,
		CreatedAt: time.Now(),
	}
	l.SetType(t)

	err := l.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return l, nil
}

//validTypes types valid
var validTypes = map[int]string{
	1: "INFO",
	2: "LOG",
	3: "WARN",
	4: "ERROR",
}

//SetType set type log
func (log *Log) SetType(t int) error {
	if tp, ok := validTypes[t]; ok {
		log.Type = tp
		return nil
	}
	return ErrNotValidTypeLog
}

//Validate validagte log
func (l *Log) Validate() error {
	if l.Type == "" || l.Text == "" {
		return ErrInvalidEntity
	}
	return nil
}
