package usecases

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type FindBookByPublisherInput struct {
	Publisher string `json:"publisher"`
}

type FindBookByPublisherOutput struct {
	Books []*entity.Book `json:"books"`
}

type IFindBookByPublisherUseCase interface {
	FindBookByPublisher(input FindBookByPublisherInput) (*FindBookByPublisherOutput, error)
}

type findBookByPublisher struct {
	bookRepository entity.IBookRepository
}

func NewFindBookByPublisherUseCase(bookRepository entity.IBookRepository) IFindBookByPublisherUseCase {
	return &findBookByPublisher{
		bookRepository: bookRepository,
	}
}

func (f *findBookByPublisher) FindBookByPublisher(input FindBookByPublisherInput) (*FindBookByPublisherOutput, error) {
	books, err := f.bookRepository.FindByPublisher(input.Publisher)
	if err != nil {
		return nil, err
	}

	return &FindBookByPublisherOutput{
		Books: books,
	}, nil
}
