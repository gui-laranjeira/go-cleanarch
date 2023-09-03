package usecases

import (
	"time"

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
	ID        string `json:"id_book"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int    `json:"pages"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
	book, err := u.bookRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	newBook, err := entity.NewBookFactory(input.Title, input.Author, input.Pages, input.Publisher, input.Year, input.ISBN)
	if err != nil {
		return nil, err
	}

	newBook.ID = book.ID
	newBook.UpdatedAt = time.Now()

	err = u.bookRepository.Update(book, newBook)
	if err != nil {
		return nil, err
	}

	return &UpdateBookOutput{
		ID:        newBook.ID.String(),
		Title:     newBook.Title,
		Author:    newBook.Author,
		Pages:     newBook.Pages,
		Publisher: newBook.Publisher,
		Year:      newBook.Year,
		ISBN:      newBook.ISBN,
		CreatedAt: newBook.CreatedAt.String(),
		UpdatedAt: newBook.UpdatedAt.String(),
	}, nil
}
