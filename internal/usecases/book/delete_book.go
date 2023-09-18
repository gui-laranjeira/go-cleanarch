package usecases

import (
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type DeleteBookInput struct {
	ID string `json:"id_book"`
}

type DeleteBookOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IDeleteBookUseCase interface {
	DeleteBook(input DeleteBookInput) (*DeleteBookOutput, error)
}

type deleteBookUseCase struct {
	bookRepository entity.IBookRepository
}

func NewDeleteBookUseCase(bookRepository entity.IBookRepository) IDeleteBookUseCase {
	return &deleteBookUseCase{
		bookRepository: bookRepository,
	}
}

func (u *deleteBookUseCase) DeleteBook(input DeleteBookInput) (*DeleteBookOutput, error) {
	rows, err := u.bookRepository.Delete(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteBookOutput{
		RowsAffected: rows,
	}, nil
}
