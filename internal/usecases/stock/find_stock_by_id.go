package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindStockByIDInput struct {
	StockID string `json:"id_stock_entry"`
}

type FindStockByIDOutput struct {
	StockID   string `json:"id_stock_entry"`
	BookID    string `json:"id_book"`
	Available bool   `json:"available"`
}

type IFindStocKByIDUseCase interface {
	FindStockByID(input FindStockByIDInput) (*FindStockByIDOutput, error)
}

type findStockByIDUseCase struct {
	stockRepository entity.IStockRepository
}

func NewFindStockByIDUseCase(stockRepository entity.IStockRepository) IFindStocKByIDUseCase {
	return &findStockByIDUseCase{
		stockRepository: stockRepository,
	}
}

func (u *findStockByIDUseCase) FindStockByID(input FindStockByIDInput) (*FindStockByIDOutput, error) {
	stockEntry, err := u.stockRepository.GetStockEntryByID(input.StockID)
	if err != nil {
		return nil, err
	}

	return &FindStockByIDOutput{
		StockID:   stockEntry.StockID.String(),
		BookID:    stockEntry.BookID.String(),
		Available: stockEntry.Available,
	}, nil
}
