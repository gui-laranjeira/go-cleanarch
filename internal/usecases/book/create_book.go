package usecases

import (
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type CreateBookInput struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int    `json:"pages"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
}

type CreateBookOutput struct {
	ID        string    `json:"id_book"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Pages     int       `json:"pages"`
	Publisher string    `json:"publisher"`
	Year      int       `json:"year"`
	ISBN      string    `json:"isbn"`
	CreatedAt time.Time `json:"created_at"`
}

type ICreateBookUseCase interface {
	CreateBook(input CreateBookInput) (*CreateBookOutput, error)
}

type createBookUseCase struct {
	bookRepository entity.IBookRepository
}

func NewCreateBookUseCase(bookRepository entity.IBookRepository) ICreateBookUseCase {
	return &createBookUseCase{
		bookRepository: bookRepository,
	}
}

func (c createBookUseCase) CreateBook(input CreateBookInput) (*CreateBookOutput, error) {
	book, err := entity.NewBookFactory(input.Title, input.Author, input.Pages, input.Publisher, input.Year, input.ISBN)
	if err != nil {
		return nil, err
	}

	err = c.bookRepository.Create(book)
	if err != nil {
		return nil, err
	}

	return &CreateBookOutput{
		ID:        book.ID.String(),
		Title:     book.Title,
		Author:    book.Author,
		Pages:     book.Pages,
		Publisher: book.Publisher,
		Year:      book.Year,
		ISBN:      book.ISBN,
		CreatedAt: book.CreatedAt,
	}, nil
}
