package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindCostumerByIDInput struct {
	ID string `json:"ID"`
}

type FindCostumerByIDOutput struct {
	ID            string `json:"id_costumer"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Document      string `json:"document"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	CurrentBookID string `json:"current_book_id"`
}

type IFindCostumerByIDUseCase interface {
	FindCostumerByID(input FindCostumerByIDInput) (*FindCostumerByIDOutput, error)
}

type findCostumerByIDUseCase struct {
	costumerRepository entity.ICostumerRepository
}

func NewFindCostumerByIDUseCase(costumerRepository entity.ICostumerRepository) IFindCostumerByIDUseCase {
	return &findCostumerByIDUseCase{
		costumerRepository: costumerRepository,
	}
}

func (c findCostumerByIDUseCase) FindCostumerByID(input FindCostumerByIDInput) (*FindCostumerByIDOutput, error) {
	costumer, err := c.costumerRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	var output FindCostumerByIDOutput

	output.ID = costumer.ID.String()
	output.Email = costumer.Email
	output.Phone = costumer.Phone
	output.Address = costumer.Address
	output.Document = costumer.Document
	output.FirstName = costumer.FirstName
	output.LastName = costumer.LastName
	output.CreatedAt = costumer.CreatedAt.String()
	output.UpdatedAt = costumer.UpdatedAt.String()
	if costumer.CurrentBookID == nil {
		output.CurrentBookID = ""
	} else {
		output.CurrentBookID = costumer.CurrentBookID.String()
	}

	return &output, nil
}
