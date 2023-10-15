package entity_test

import (
	"testing"
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewCostumer(t *testing.T) {
	type test struct {
		email      string
		phone      string
		address    string
		document   string
		first_name string
		last_name  string
		created_at time.Time
		expect     error
	}

	tests := []test{
		{
			email:      "teste@teste.com",
			phone:      "11999999999",
			address:    "R. de Teste",
			document:   "12345678",
			first_name: "Maria",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     nil,
		},
		{
			email:      "",
			phone:      "11999999999",
			address:    "R. de Teste",
			document:   "12345678",
			first_name: "Maria",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			phone:      "",
			address:    "R. de Teste",
			document:   "12345678",
			first_name: "Maria",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			phone:      "11999999999",
			address:    "",
			document:   "12345678",
			first_name: "Maria",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			phone:      "11999999999",
			address:    "R. de Teste",
			document:   "",
			first_name: "Maria",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			phone:      "11999999999",
			address:    "R. de Teste",
			document:   "12345678",
			first_name: "",
			last_name:  "Teste",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
		{
			email:      "teste@teste.com",
			phone:      "11999999999",
			address:    "R. de Teste",
			document:   "12345678",
			first_name: "Maria",
			last_name:  "",
			created_at: time.Now(),
			expect:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewCostumerFactory(tc.email, tc.phone, tc.address, tc.document, tc.first_name, tc.last_name)
		assert.Equal(t, tc.expect, err)
	}
}
