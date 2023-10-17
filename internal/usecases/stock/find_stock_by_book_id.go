package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindStockByBookIDInput struct {
	BookID string `json:"id_book"`
}

type FindStockByBookIDOutput struct {
	Stock []*entity.StockEntry
}

type IFindStockByBookIDUseCase interface {
	FindStockByBookID(input FindStockByBookIDInput) (*FindStockByBookIDOutput, error)
}

type findStockByBookIDUseCase struct {
	stockRepository entity.IStockRepository
}

func NewFindStockByBookIDUseCase(stockRepository entity.IStockRepository) IFindStockByBookIDUseCase {
	return &findStockByBookIDUseCase{
		stockRepository: stockRepository,
	}
}

func (u *findStockByBookIDUseCase) FindStockByBookID(input FindStockByBookIDInput) (*FindStockByBookIDOutput, error) {
	stockEntry, err := u.stockRepository.GetStockEntryByBookID(input.BookID)
	if err != nil {
		return nil, err
	}

	return &FindStockByBookIDOutput{
		stockEntry,
	}, nil
}
