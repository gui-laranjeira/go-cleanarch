package usecases

import (
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type FindUserByIDInput struct {
	ID string `json:"ID"`
}

type FindUserByIDOutput struct {
	ID        string `json:"id_user"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IFindUserByIDUseCase interface {
	FindUserByID(input FindUserByIDInput) (*FindUserByIDOutput, error)
}

type findUserByIDUseCase struct {
	userRepository entity.IUserRepository
}

func NewFindUserByIDUseCase(repo entity.IUserRepository) IFindUserByIDUseCase {
	return &findUserByIDUseCase{
		userRepository: repo,
	}
}

func (c findUserByIDUseCase) FindUserByID(input FindUserByIDInput) (*FindUserByIDOutput, error) {
	user, err := c.userRepository.GetUserByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindUserByIDOutput{
		ID:        user.ID.String(),
		Email:     user.Email,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
