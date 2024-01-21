package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStockEntry(t *testing.T) {
	r, err := entity.NewStockEntryFactory(uuid.New().String(), true)
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestValidateStock(t *testing.T) {
	type test struct {
		id_book   string
		available bool
		expect    error
	}

	tests := []test{
		{
			id_book:   uuid.New().String(),
			available: true,
			expect:    nil,
		},
		{
			id_book:   "",
			available: true,
			expect:    entity.ErrInvalidEntity,
		},
		{
			id_book:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			available: true,
			expect:    entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewStockEntryFactory(tc.id_book, tc.available)
		assert.Equal(t, tc.expect, err)
	}
}
