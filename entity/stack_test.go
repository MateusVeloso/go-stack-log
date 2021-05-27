package entity_test

import (
	"testing"

	"github.com/mateusveloso/go-stack-log/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s, err := entity.NewStack("Cadastro de usuário.")
	assert.Nil(t, err)
	assert.Equal(t, s.Title, "Cadastro de usuário.")
	assert.NotNil(t, s.ID)
	assert.NotNil(t, s.CreatedAt)
}

func TestAddLog(t *testing.T) {
	s, _ := entity.NewStack("Cadastro de usuário.")
	err := s.RemoveLog(entity.NewID())
	assert.Equal(t, entity.ErrNotFound, err)
	logID := entity.NewID()
	_ = s.AddLog(logID)
	err = s.RemoveLog(logID)
	assert.Nil(t, err)
}

func TestGetLog(t *testing.T) {
	s, _ := entity.NewStack("Cadastro de usuário.")
	logID := entity.NewID()
	_ = s.AddLog(logID)
	id, err := s.GetLog(logID)
	assert.Nil(t, err)
	assert.Equal(t, id, logID)
	_, err = s.GetLog(entity.NewID())
	assert.Equal(t, entity.ErrNotFound, err)
}

func TestStackValidate(t *testing.T) {
	type test struct {
		title string
		want  error
	}

	tests := []test{
		{
			title: "Cadastro de usuário",
			want:  nil,
		},
		{
			title: "",
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewStack(tc.title)
		assert.Equal(t, err, tc.want)
	}

}
