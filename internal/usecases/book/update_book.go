package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type UpdateBookInput struct {
	ID        string `json:"id_book"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int    `json:"pages"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
}

type UpdateBookOutput struct {
	RowsAffected int64 `json:"rows_affected"`
}

type IUpdateBookUseCase interface {
	UpdateBook(input UpdateBookInput) (*UpdateBookOutput, error)
}

type updateBookUseCase struct {
	bookRepository entity.IBookRepository
}

func NewUpdateBookUseCase(bookRepository entity.IBookRepository) IUpdateBookUseCase {
	return &updateBookUseCase{
		bookRepository: bookRepository,
	}
}

func (u *updateBookUseCase) UpdateBook(input UpdateBookInput) (*UpdateBookOutput, error) {
	newBook, err := entity.NewBookFactory(input.Title, input.Author, input.Pages, input.Publisher, input.Year, input.ISBN)
	if err != nil {
		return nil, err
	}

	id := uuid.MustParse(input.ID)

	newBook.ID = id
	newBook.UpdatedAt = time.Now()

	//TODO implement rows affected treatment
	rows, err := u.bookRepository.Update(newBook)
	if err != nil {
		return nil, err
	}

	return &UpdateBookOutput{
		RowsAffected: rows,
	}, nil
}
