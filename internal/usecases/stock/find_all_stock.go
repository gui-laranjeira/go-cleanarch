package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindAllStockOutput struct {
	Stock []*entity.StockEntry `json:"stock"`
}

type IFindAllStockUseCase interface {
	FindAllStock() (*FindAllStockOutput, error)
}

type findAllStockUseCase struct {
	stockRepository entity.IStockRepository
}

func NewFindAllStockUseCase(stockRepository entity.IStockRepository) IFindAllStockUseCase {
	return &findAllStockUseCase{
		stockRepository: stockRepository,
	}
}

func (u *findAllStockUseCase) FindAllStock() (*FindAllStockOutput, error) {
	stock, err := u.stockRepository.GetAllStock()
	if err != nil {
		return nil, err
	}

	return &FindAllStockOutput{
		stock,
	}, nil
}
