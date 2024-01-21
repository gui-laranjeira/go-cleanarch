package entity

import (
	"github.com/google/uuid"
)

type StockEntry struct {
	BookID    uuid.UUID `json:"id_book"`
	StockID   uuid.UUID `json:"id_stock_entry"`
	Available bool      `json:"available"`
}

type IStockRepository interface {
	CreateStockEntry(entry *StockEntry) error
	GetAllStock() ([]*StockEntry, error)
	GetStockEntryByID(id string) (*StockEntry, error)
	GetStockEntryByBookID(id_book string) ([]*StockEntry, error)
	DeleteStockEntry(id_stock_entry string) (int64, error)
	BorrowBook(id_stock_entry string) error
	ReturnBook(id_stock_entry string) error
}

func NewStockEntryFactory(book_id string, available bool) (*StockEntry, error) {
	bookId, err := uuid.Parse(book_id)
	if err != nil {
		return nil, ErrInvalidEntity
	}

	r := StockEntry{
		BookID:    bookId,
		StockID:   uuid.New(),
		Available: available,
	}

	return &r, nil
}
