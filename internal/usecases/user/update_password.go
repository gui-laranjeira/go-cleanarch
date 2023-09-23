package usecases

import (
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdatePasswordInput struct {
	ID       string
	Password string
}

type UpdatePasswordOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IUpdatePasswordUseCase interface {
	UpdatePassword(UpdatePasswordInput) (*UpdatePasswordOutput, error)
}

type updatePasswordUseCase struct {
	userRepository entity.IUserRepository
}

func NewUpdatePasswordUseCase(userRepository entity.IUserRepository) IUpdatePasswordUseCase {
	return &updatePasswordUseCase{
		userRepository: userRepository,
	}
}

func (u *updatePasswordUseCase) UpdatePassword(input UpdatePasswordInput) (*UpdatePasswordOutput, error) {
	rows, err := u.userRepository.ChangePassword(input.ID, input.Password)
	if err != nil {
		return nil, err
	}

	return &UpdatePasswordOutput{
		RowsAffected: rows,
	}, nil
}
