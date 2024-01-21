package entity_test

import (
	"testing"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBorrow(t *testing.T) {
	b, err := entity.NewBorrowEntryFactory("6eee0146-8d92-4a14-9123-842c98832a7c", "6eee0146-8d92-4a14-9123-842c98832a7c")
	assert.Nil(t, err)
	assert.NotNil(t, b)
}

func TestValidateBorrow(t *testing.T) {
	type test struct {
		id_costumer    string
		id_stock_entry string
		expect         error
	}

	tests := []test{
		{
			id_costumer:    "6eee0146-8d92-4a14-9123-842c98832a7c",
			id_stock_entry: "6eee0146-8d92-4a14-9123-842c98832a7c",
			expect:         nil,
		},
		{
			id_costumer:    "",
			id_stock_entry: "6eee0146-8d92-4a14-9123-842c98832a7c",
			expect:         entity.ErrInvalidEntity,
		},
		{
			id_costumer:    "6eee0146-8d92-4a14-9123-842c98832a7c",
			id_stock_entry: "",
			expect:         entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Logf(`Costumer "%s"`, tc.id_costumer)
		t.Logf(`Stock "%s"`, tc.id_stock_entry)
		_, err := entity.NewBorrowEntryFactory(tc.id_costumer, tc.id_stock_entry)
		assert.Equal(t, tc.expect, err)
	}

}
