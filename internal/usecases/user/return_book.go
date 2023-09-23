package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type ReturnBookUsecaseInput struct {
	UserID string
	BookID string
}

type IReturnBookUsecase interface {
	ReturnBook(ReturnBookUsecaseInput) error
}

type ReturnBookUsecase struct {
	userRepository entity.IUserRepository
}

func NewReturnBookUseCase(userRepository entity.IUserRepository) IReturnBookUsecase {
	return &ReturnBookUsecase{
		userRepository: userRepository,
	}
}

func (u *ReturnBookUsecase) ReturnBook(input ReturnBookUsecaseInput) error {
	err := u.userRepository.ReturnBook(input.UserID, input.BookID)
	if err != nil {
		return err
	}
	return nil
}
