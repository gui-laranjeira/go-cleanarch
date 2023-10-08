package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewStockEntry(t *testing.T) {
	r, err := entity.NewStockEntryFactory(uuid.New(), "S")
	assert.Nil(t, err)
	assert.NotNil(t, r)
}
