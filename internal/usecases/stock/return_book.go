package usecases

import (
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type ReturnBookInput struct {
	StockEntryID string `json:"id_stock_entry"`
	CostumerID   string `json:"id_costumer"`
}

type ReturnBookOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IReturnBookUseCase interface {
	ReturnBook(entry ReturnBookInput) (*ReturnBookOutput, error)
}

type returnBookUseCase struct {
	borrowEntryRepository entity.IBorrowEntryRepository
	stockRepository       entity.IStockRepository
}

func NewReturnBookUseCase(borrowEntryRepository entity.IBorrowEntryRepository, stockRepository entity.IStockRepository) IReturnBookUseCase {
	return &returnBookUseCase{
		borrowEntryRepository: borrowEntryRepository,
		stockRepository:       stockRepository,
	}
}

func (u *returnBookUseCase) ReturnBook(input ReturnBookInput) (*ReturnBookOutput, error) {
	borrowEntry, err := u.borrowEntryRepository.GetBorrowEntryByID(input.StockEntryID, input.CostumerID)
	if err != nil {
		return nil, err
	}

	borrowEntry.Returned = true
	borrowEntry.UpdatedAt = time.Now()

	err = u.stockRepository.ReturnBook(input.StockEntryID)
	if err != nil {
		return nil, err
	}

	rows, err := u.borrowEntryRepository.Return(borrowEntry)
	if err != nil {
		return nil, err
	}
	return &ReturnBookOutput{RowsAffected: rows}, nil
}
