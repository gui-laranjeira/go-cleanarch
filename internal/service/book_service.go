package service

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

type bookService struct {
	repo entity.IBookRepository
}

func NewBookService(repo entity.IBookRepository) entity.IBookUseCase {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) CreateBook(book *entity.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *entity.Book, newBook *entity.Book) error {
	return s.repo.Update(book, newBook)
}

func (s *bookService) GetBooks() ([]*entity.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id string) (*entity.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) GetBookByTitle(title string) (*entity.Book, error) {
	return s.repo.FindByTitle(title)
}

func (s *bookService) GetBookByAuthor(author string) ([]*entity.Book, error) {
	return s.repo.FindByAuthor(author)
}

func (s *bookService) GetBookByPublisher(publisher string) ([]*entity.Book, error) {
	return s.repo.FindByPublisher(publisher)
}
