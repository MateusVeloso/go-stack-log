package entity_test

import (
	"testing"

	"github.com/mateusveloso/go-stack-log/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	text := "Erro ao criar usuário, coluna ID não existe."
	b, err := entity.NewLog(1, text)
	assert.Nil(t, err)
	assert.Equal(t, b.Text, text)
	assert.Equal(t, b.Type, "INFO")
	assert.NotNil(t, b.ID)
	assert.NotNil(t, b.CreatedAt)
}

func TestLogValidate(t *testing.T) {
	type test struct {
		Type int
		Text string
		want error
	}

	tests := []test{
		{
			Type: 1,
			Text: "Erro ao criar usuário, coluna ID não existe.",
			want: nil,
		},
		{
			Type: 1,
			Text: "",
			want: entity.ErrInvalidEntity,
		},
		{
			Type: 0,
			Text: "Erro ao criar usuário, coluna ID não existe.",
			want: entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewLog(tc.Type, tc.Text)
		assert.Equal(t, err, tc.want)
	}

}
