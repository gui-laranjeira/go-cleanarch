package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockEntry struct {
	Book_ID    uuid.UUID `json:"id_book"`
	User_ID    uuid.UUID `json:"id_user"`
	Avaiable   bool      `json:"avaiable"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

func NewStockEntry(book_id uuid.UUID, user_id uuid.UUID, avaiable bool) (*StockEntry, error) {
	r := StockEntry{
		Book_ID:  book_id,
		User_ID:  user_id,
		Avaiable: avaiable,
	}

	err := r.validateStockEntry()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (s *StockEntry) validateStockEntry() error {
	if s.Book_ID.String() == "" || s.User_ID.String() == "" {
		return ErrInvalidEntity
	}
	return nil
}
