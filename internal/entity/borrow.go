package entity

import (
	"time"

	"github.com/google/uuid"
)

type BorrowEntry struct {
	CostumerID   uuid.UUID `json:"id_costumer"`
	StockEntryID uuid.UUID `json:"id_stock_entry"`
	BorrowDate   time.Time `json:"borrow_date"`
	DueDate      time.Time `json:"due_date"`
	Returned     bool      `json:"return_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IBorrowEntryRepository interface {
	Borrow(entry *BorrowEntry) (int64, error)
	Return(entry *BorrowEntry) (int64, error)
	GetBorrowEntryByID(id_stock_entry string, id_costumer string) (*BorrowEntry, error)
}

type IBorrowEntryUseCase interface {
	BorrowBook(entry *BorrowEntry) error
	ReturnBook(entry *BorrowEntry) error
}

func NewBorrowEntryFactory(id_costumer string, id_stock_entry string) (*BorrowEntry, error) {
	costumerId, _ := uuid.Parse(id_costumer)
	stockEntryID, _ := uuid.Parse(id_stock_entry)

	r := BorrowEntry{
		CostumerID:   costumerId,
		StockEntryID: stockEntryID,
		BorrowDate:   time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 15),
		Returned:     false,
		CreatedAt:    time.Now(),
	}
	err := r.ValidateBorrowEntry()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return &r, nil
}

func (b *BorrowEntry) ValidateBorrowEntry() error {
	if b.CostumerID.String() == "" || b.StockEntryID.String() == "" || b.BorrowDate.String() == "" || b.DueDate.String() == "" || b.CostumerID.String() == "00000000-0000-0000-0000-000000000000" || b.StockEntryID.String() == "00000000-0000-0000-0000-000000000000" {
		return ErrInvalidEntity
	}
	return nil
}
