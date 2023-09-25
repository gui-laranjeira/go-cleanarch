package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdateCostumerInput struct {
	ID        string
	Email     string
	Phone     string
	Adress    string
	Document  string
	FirstName string
	LastName  string
}

type UpdateCostumerOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IUpdateCostumerUseCase interface {
	UpdateCostumer(UpdateCostumerInput) (*UpdateCostumerOutput, error)
}

type updateCostumeruseCase struct {
	costumerRepository entity.ICostumerRepository
}

func NewUpdateCostumerUseCase(CostumerRepository entity.ICostumerRepository) IUpdateCostumerUseCase {
	return &updateCostumeruseCase{
		costumerRepository: CostumerRepository,
	}
}

func (u *updateCostumeruseCase) UpdateCostumer(input UpdateCostumerInput) (*UpdateCostumerOutput, error) {
	newCostumer, err := entity.NewCostumerFactory(input.Email, input.Phone, input.Adress, input.Document, input.FirstName, input.LastName)
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
