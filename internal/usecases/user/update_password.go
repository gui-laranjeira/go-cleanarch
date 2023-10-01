package usecases

import (
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordInput struct {
	ID       string `json:"id_user"`
	Password string `json:"password"`
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
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, entity.ErrPasswordHashing
	}

	rows, err := u.userRepository.ChangePassword(input.ID, string(hash))
	if err != nil {
		return nil, err
	}

	return &UpdatePasswordOutput{
		RowsAffected: rows,
	}, nil
}
