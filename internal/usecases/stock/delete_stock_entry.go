package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type DeleteStockEntryInput struct {
	StockID string `json:"id_stock_entry"`
}

type DeleteStockEntryOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IDeleteStockEntryUseCase interface {
	DeleteStockEntry(DeleteStockEntryInput) (*DeleteStockEntryOutput, error)
}

type deleteStockEntryUseCase struct {
	stockRepository entity.IStockRepository
}

func NewDeleteStockEntryUseCase(stockRepository entity.IStockRepository) IDeleteStockEntryUseCase {
	return &deleteStockEntryUseCase{
		stockRepository: stockRepository,
	}
}

func (u *deleteStockEntryUseCase) DeleteStockEntry(input DeleteStockEntryInput) (*DeleteStockEntryOutput, error) {
	rows, err := u.stockRepository.DeleteStockEntry(input.StockID)
	if err != nil {
		return nil, err
	}

	return &DeleteStockEntryOutput{
		rows,
	}, err
}
