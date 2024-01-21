package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type BorrowBookInput struct {
	StockEntryID string `json:"id_stock_entry"`
	CostumerID   string `json:"id_costumer"`
}

type BorrowBookOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IBorrowBookUseCase interface {
	BorrowBook(entry BorrowBookInput) (*BorrowBookOutput, error)
}

type borrowBookUseCase struct {
	borrowEntryRepository entity.IBorrowEntryRepository
	stockRepository       entity.IStockRepository
}

func NewBorrowBookUseCase(borrowEntryRepository entity.IBorrowEntryRepository, stockRepository entity.IStockRepository) IBorrowBookUseCase {
	return &borrowBookUseCase{
		borrowEntryRepository: borrowEntryRepository,
		stockRepository:       stockRepository,
	}
}

func (u *borrowBookUseCase) BorrowBook(entry BorrowBookInput) (*BorrowBookOutput, error) {
	borrowEntry, err := entity.NewBorrowEntryFactory(entry.CostumerID, entry.StockEntryID)
	if err != nil {
		return nil, err
	}

	stockEntry, err := u.stockRepository.GetStockEntryByID(entry.StockEntryID)
	if err != nil {
		return nil, err
	}
	if stockEntry.Available == false {
		return nil, err
	}

	err = u.stockRepository.BorrowBook(stockEntry.StockID.String())
	if err != nil {
		return nil, err
	}

	rows, err := u.borrowEntryRepository.Borrow(borrowEntry)
	if err != nil {
		return nil, err
	}
	return &BorrowBookOutput{RowsAffected: rows}, nil
}
