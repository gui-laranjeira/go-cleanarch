package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type ReturnBookUsecaseInput struct {
	CostumerID string
	BookID     string
}

type IReturnBookUsecase interface {
	ReturnBook(ReturnBookUsecaseInput) error
}

type returnBookUsecase struct {
	costumerRepository entity.ICostumerRepository
}

func NewReturnBookUseCase(costumerRepository entity.ICostumerRepository) IReturnBookUsecase {
	return &returnBookUsecase{
		costumerRepository: costumerRepository,
	}
}

func (u *returnBookUsecase) ReturnBook(input ReturnBookUsecaseInput) error {
	err := u.costumerRepository.ReturnBook(input.CostumerID, input.BookID)
	if err != nil {
		return err
	}
	return nil
}
