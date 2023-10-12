package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindBookByIDInput struct {
	ID string `json:"ID"`
}

type FindBookByIDOutput struct {
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

type IFindBookByIDUseCase interface {
	FindBookByID(input FindBookByIDInput) (*FindBookByIDOutput, error)
}

type findBookByID struct {
	bookRepository entity.IBookRepository
}

func NewFindBookByIDUseCase(bookRepository entity.IBookRepository) IFindBookByIDUseCase {
	return &findBookByID{
		bookRepository: bookRepository,
	}
}

func (f *findBookByID) FindBookByID(input FindBookByIDInput) (*FindBookByIDOutput, error) {
	book, err := f.bookRepository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &FindBookByIDOutput{
		ID:        book.ID.String(),
		Title:     book.Title,
		Author:    book.Author,
		Pages:     book.Pages,
		Publisher: book.Publisher,
		Year:      book.Year,
		ISBN:      book.ISBN,
		CreatedAt: book.CreatedAt.String(),
		UpdatedAt: book.UpdatedAt.String(),
	}, nil
}
