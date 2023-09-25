package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type CreateCostumerInput struct {
	Email     string
	Phone     string
	Adress    string
	Document  string
	Password  string
	FirstName string
	LastName  string
}

type CreateCostumerOutput struct {
	Error error
}

type ICreateCostumerUseCase interface {
	CreateCostumer(input CreateCostumerInput) error
}

type createCostumerUseCase struct {
	costumerRepository entity.ICostumerRepository
}

func NewCreateCostumerUseCase(CostumerRepository entity.ICostumerRepository) ICreateCostumerUseCase {
	return &createCostumerUseCase{
		costumerRepository: CostumerRepository,
	}
}

func (c *createCostumerUseCase) CreateCostumer(input CreateCostumerInput) error {
	costumer, err := entity.NewCostumerFactory(input.Email, input.Phone, input.Adress, input.Document, input.FirstName, input.LastName)
	if err != nil {
		return nil
	}

	err = c.costumerRepository.Create(costumer)
	if err != nil {
		return nil
	}

	return nil
}
