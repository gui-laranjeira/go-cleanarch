package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindBookByTitleInput struct {
	Title string `json:"title"`
}

type FindBookByTitleOutput struct {
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

type IFindBookByTitleUseCase interface {
	FindBookByTitle(input FindBookByTitleInput) (*FindBookByTitleOutput, error)
}

type findbookByTitle struct {
	bookRepository entity.IBookRepository
}

func NewFindBookByTitleUseCase(bookRepository entity.IBookRepository) IFindBookByTitleUseCase {
	return &findbookByTitle{
		bookRepository: bookRepository,
	}
}

func (f *findbookByTitle) FindBookByTitle(input FindBookByTitleInput) (*FindBookByTitleOutput, error) {
	book, err := f.bookRepository.FindByTitle(input.Title)
	if err != nil {
		return nil, err
	}

	return &FindBookByTitleOutput{
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
