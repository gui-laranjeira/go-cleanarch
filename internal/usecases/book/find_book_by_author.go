package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindBookByAuthorInput struct {
	Author string `json:"author"`
}

type FindBookByAuthorOutput struct {
	Books []*entity.Book `json:"books"`
}

type IFindBookByAuthorUseCase interface {
	FindBookByAuthor(input FindBookByAuthorInput) (*FindBookByAuthorOutput, error)
}

type findBookByAuthor struct {
	bookRepository entity.IBookRepository
}

func NewFindBookByAuthorUseCase(bookRepository entity.IBookRepository) IFindBookByAuthorUseCase {
	return &findBookByAuthor{
		bookRepository: bookRepository,
	}
}

func (f findBookByAuthor) FindBookByAuthor(input FindBookByAuthorInput) (*FindBookByAuthorOutput, error) {
	books, err := f.bookRepository.FindByAuthor(input.Author)
	if err != nil {
		return nil, err
	}

	return &FindBookByAuthorOutput{
		Books: books,
	}, nil
}
