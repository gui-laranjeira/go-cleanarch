package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type CreateUserInput struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
		return err
	}

	err = c.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
