package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdateUserInput struct {
	ID        string
	Email     string
	Phone     string
	FirstName string
	LastName  string
}

type UpdateUserOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IUpdateUserUseCase interface {
	UpdateUser(UpdateUserInput) (*UpdateUserOutput, error)
}

type updateUseruseCase struct {
	userRepository entity.IUserRepository
}

func NewUpdateUserUseCase(userRepository entity.IUserRepository) IUpdateUserUseCase {
	return &updateUseruseCase{
		userRepository: userRepository,
	}
}

func (u *updateUseruseCase) UpdateUser(input UpdateUserInput) (*UpdateUserOutput, error) {
	newUser, err := entity.NewUserFactory(input.Email, input.Phone, "", input.FirstName, input.LastName)
	if err != nil {
		return nil, err
	}

	id := uuid.MustParse(input.ID)

	newUser.ID = id
	newUser.UpdatedAt = time.Now()

	rows, err := u.userRepository.Update(newUser)
	if err != nil {
		return nil, err
	}

	return &UpdateUserOutput{
		RowsAffected: rows,
	}, nil
}
