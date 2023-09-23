package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type BorrowBookUsecaseInput struct {
	UserID string
	BookID string
}

type IBorrowBookUsecase interface {
	BorrowBook(BorrowBookUsecaseInput) error
}

type borrowBookUsecase struct {
	userRepository entity.IUserRepository
}

func NewBorrowBookUseCase(userRepository entity.IUserRepository) IBorrowBookUsecase {
	return &borrowBookUsecase{
		userRepository: userRepository,
	}
}

func (u *borrowBookUsecase) BorrowBook(input BorrowBookUsecaseInput) error {
	err := u.userRepository.BorrowBook(input.UserID, input.BookID)
	if err != nil {
		return err
	}
	return nil
}
