package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindAllBooksOutput struct {
	Books []*entity.Book `json:"books"`
}

type IFindAllBooksUseCase interface {
	FindAllBooks() (*FindAllBooksOutput, error)
}

type findAllBooksUseCase struct {
	bookRepository entity.IBookRepository
}

func NewFindAllBooksUseCase(bookRepository entity.IBookRepository) IFindAllBooksUseCase {
	return &findAllBooksUseCase{
		bookRepository: bookRepository,
	}
}

func (c findAllBooksUseCase) FindAllBooks() (*FindAllBooksOutput, error) {
	books, err := c.bookRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &FindAllBooksOutput{
		books,
	}, nil
}
