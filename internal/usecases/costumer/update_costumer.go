package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdateCostumerInput struct {
	ID        string `json:"id_costumer"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Document  string `json:"document"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateCostumerOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IUpdateCostumerUseCase interface {
	UpdateCostumer(UpdateCostumerInput) (*UpdateCostumerOutput, error)
}

type updateCostumerUseCase struct {
	costumerRepository entity.ICostumerRepository
}

func NewUpdateCostumerUseCase(CostumerRepository entity.ICostumerRepository) IUpdateCostumerUseCase {
	return &updateCostumerUseCase{
		costumerRepository: CostumerRepository,
	}
}

func (u *updateCostumerUseCase) UpdateCostumer(input UpdateCostumerInput) (*UpdateCostumerOutput, error) {
	newCostumer, err := entity.NewCostumerFactory(input.Email, input.Phone, input.Address, input.Document, input.FirstName, input.LastName)
	if err != nil {
		return nil, err
	}

	id := uuid.MustParse(input.ID)

	newCostumer.ID = id
	newCostumer.UpdatedAt = time.Now()

	rows, err := u.costumerRepository.Update(newCostumer)
	if err != nil {
		return nil, err
	}

	return &UpdateCostumerOutput{
		RowsAffected: rows,
	}, nil
}
