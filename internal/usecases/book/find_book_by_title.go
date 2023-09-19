package usecases

import (
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
)

type FindBookByTitleInput struct {
	Title string `json:"title"`
}

type FindBookByTitleOutput struct {
	Books []*entity.Book
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
	books, err := f.bookRepository.FindByTitle(input.Title)
	if err != nil {
		return nil, err
	}

	return &FindBookByTitleOutput{
		Books: books,
	}, nil
}
