package entity_test

import (
	"testing"
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type test struct {
		email      string
		password   string
		first_name string
		last_name  string
		created_at time.Time
		expect     error
	}
	tests := []test{
		{
			email:      "teste@teste.com",
			password:   "123456",
			first_name: "Gui",
			last_name:  "Laranjeira",
			created_at: time.Now(),
			expect:     nil,
		},
		{
			email:      "",
			password:   "123456",
			first_name: "Gui",
			last_name:  "Laranjeira",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			password:   "",
			first_name: "Gui",
			last_name:  "Laranjeira",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			password:   "123456",
			first_name: "",
			last_name:  "Laranjeira",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			password:   "123456",
			first_name: "Gui",
			last_name:  "",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			password:   "",
			first_name: "Gui",
			last_name:  "Laranjeira",
			expect:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewUserFactory(tc.email, tc.password, tc.first_name, tc.last_name)
		assert.Equal(t, tc.expect, err)
	}
}
