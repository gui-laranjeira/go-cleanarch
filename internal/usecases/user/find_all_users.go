package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindAllUsersOutput struct {
	Users []*entity.User `json:"users"`
}

type IFindAllUsersUseCase interface {
	FindAllUsers() (*FindAllUsersOutput, error)
}

type findAllUsersUseCase struct {
	userRepository entity.IUserRepository
}

func NewFindAllUsersUseCase(userRepository entity.IUserRepository) IFindAllUsersUseCase {
	return &findAllUsersUseCase{
		userRepository: userRepository,
	}
}

func (c findAllUsersUseCase) FindAllUsers() (*FindAllUsersOutput, error) {
	users, err := c.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return &FindAllUsersOutput{
		users,
	}, nil
}
