package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindAllCostumersOutput struct {
	Costumers []*entity.Costumer `json:"costumers"`
}

type IFindAllCostumersUseCase interface {
	FindAllCostumers() (*FindAllCostumersOutput, error)
}

type findAllCostumersUseCase struct {
	costumerRepository entity.ICostumerRepository
}

func NewFindAllCostumersUseCase(costumerRepository entity.ICostumerRepository) IFindAllCostumersUseCase {
	return &findAllCostumersUseCase{
		costumerRepository: costumerRepository,
	}
}

func (c findAllCostumersUseCase) FindAllCostumers() (*FindAllCostumersOutput, error) {
	costumers, err := c.costumerRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &FindAllCostumersOutput{
		costumers,
	}, nil
}
