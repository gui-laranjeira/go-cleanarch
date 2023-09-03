package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockEntry struct {
	Book_ID    uuid.UUID `json:"id_book"`
	Stock_ID   uuid.UUID `json:"id_stock"`
	User_ID    uuid.UUID `json:"id_user"`
	Avaiable   bool      `json:"avaiable"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

func NewStockEntryFactory(book_id uuid.UUID, avaiable bool) (*StockEntry, error) {
	r := StockEntry{
		Book_ID:  book_id,
		Stock_ID: uuid.New(),
		Avaiable: avaiable,
	}

	err := r.validateStockEntry()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (s *StockEntry) validateStockEntry() error {
	if s.Book_ID.String() == "" || s.Stock_ID.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
