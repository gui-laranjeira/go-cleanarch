package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdateUserInput struct {
	ID        string `json:"id_user"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
	newUser, err := entity.NewUserFactory(input.Email, input.Phone, "123456", input.FirstName, input.LastName)
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
