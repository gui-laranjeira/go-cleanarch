package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type CreateStockEntryInput struct {
	BookID string `json:"id_book"`
}

type CreateStockEntryOutput struct {
	StockID   string `json:"id_stock_entry"`
	BookID    string `json:"id_book"`
	Available bool   `json:"available"`
}

type ICreateStockEntryUseCase interface {
	CreateStockEntry(input CreateStockEntryInput) (*CreateStockEntryOutput, error)
}

type createStockEntryUseCase struct {
	stockRepository entity.IStockRepository
}

func NewCreateStockEntryUseCase(stockRepository entity.IStockRepository) ICreateStockEntryUseCase {
	return &createStockEntryUseCase{
		stockRepository: stockRepository,
	}
}

func (u *createStockEntryUseCase) CreateStockEntry(input CreateStockEntryInput) (*CreateStockEntryOutput, error) {
	stockEntry, err := entity.NewStockEntryFactory(input.BookID, true)
	if err != nil {
		return nil, err
	}

	err = u.stockRepository.CreateStockEntry(stockEntry)
	if err != nil {
		return nil, err
	}

	return &CreateStockEntryOutput{
		StockID:   stockEntry.StockID.String(),
		BookID:    stockEntry.BookID.String(),
		Available: stockEntry.Available,
	}, nil
}
