package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type BorrowBookUsecaseInput struct {
	CostumerID string
	BookID     string
}

type IBorrowBookUsecase interface {
	BorrowBook(BorrowBookUsecaseInput) error
}

type borrowBookUsecase struct {
	costumerRepository entity.ICostumerRepository
}

func NewBorrowBookUseCase(CostumerRepository entity.ICostumerRepository) IBorrowBookUsecase {
	return &borrowBookUsecase{
		costumerRepository: CostumerRepository,
	}
}

func (u *borrowBookUsecase) BorrowBook(input BorrowBookUsecaseInput) error {
	err := u.costumerRepository.BorrowBook(input.CostumerID, input.BookID)
	if err != nil {
		return err
	}
	return nil
}
