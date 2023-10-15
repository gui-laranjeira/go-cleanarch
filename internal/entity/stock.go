package entity

import (
	"github.com/google/uuid"
)

type StockEntry struct {
	BookID     uuid.UUID `json:"id_book"`
	StockID    uuid.UUID `json:"id_stock"`
	Available  string    `json:"available"`
	CostumerID uuid.UUID `json:"id_user"`
}

type IStockRepository interface {
	CreateStockEntry(entry *StockEntry) error
	BorrowBook(entry *StockEntry, costumer *Costumer) error
	ReturnBook(entry *StockEntry, costumer *Costumer) error
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
