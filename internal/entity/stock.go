package entity

import (
	"github.com/google/uuid"
)

type StockEntry struct {
	BookID    uuid.UUID `json:"id_book"`
	StockID   uuid.UUID `json:"id_stock"`
	Available string    `json:"avaiable"`
}

func NewStockEntryFactory(book_id uuid.UUID, available string) (*StockEntry, error) {
	r := StockEntry{
		BookID:    book_id,
		StockID:   uuid.New(),
		Available: available,
	}

	err := r.validateStockEntry()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (s *StockEntry) validateStockEntry() error {
	if s.BookID.String() == "" || s.StockID.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
