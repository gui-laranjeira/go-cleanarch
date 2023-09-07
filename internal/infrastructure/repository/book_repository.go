package repository

import "github.com/gui-laranjeira/go-cleanarch/internal/entity"

func Create(book *entity.Book) error {
	return nil
}

func Update(book *entity.Book, newBook *entity.Book) error {
	return nil
}

func FindAll() ([]*entity.Book, error) {
	return nil, nil
}

func FindByID(id string) (*entity.Book, error) {
	return nil, nil
}

func FindByTitle(title string) (*entity.Book, error) {
	return nil, nil
}

func FindByAuthor(author string) ([]*entity.Book, error) {
	return nil, nil
}

func FindByPublisher(publisher string) ([]*entity.Book, error) {
	return nil, nil
}
