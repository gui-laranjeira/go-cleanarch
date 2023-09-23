package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type CreateUserInput struct {
	Email     string
	Phone     string
	Password  string
	FirstName string
	LastName  string
}

type CreateUserOutput struct {
	Error error
}

type ICreateUserUseCase interface {
	CreateUser(input CreateUserInput) error
}

type createUserUseCase struct {
	userRepository entity.IUserRepository
}

func NewCreateUserUseCase(userRepository entity.IUserRepository) ICreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}

func (c *createUserUseCase) CreateUser(input CreateUserInput) error {
	user, err := entity.NewUserFactory(input.Email, input.Phone, input.Password, input.FirstName, input.LastName)
	if err != nil {
		return nil
	}

	err = c.userRepository.Create(user)
	if err != nil {
		return nil
	}

	return nil
}
