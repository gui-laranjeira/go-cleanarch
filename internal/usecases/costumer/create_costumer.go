package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type CreateCostumerInput struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Document  string `json:"document"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
	costumer, err := entity.NewCostumerFactory(input.Email, input.Phone, input.Address, input.Document, input.FirstName, input.LastName)
	if err != nil {
		return nil
	}

	err = c.costumerRepository.Create(costumer)
	if err != nil {
		return err
	}

	return nil
}
