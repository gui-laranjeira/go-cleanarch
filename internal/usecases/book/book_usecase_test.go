package usecases_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/book"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) Create(book *entity.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Update(newBook *entity.Book) (int64, error) {
	args := m.Called(newBook)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBookRepository) FindAll() ([]*entity.Book, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Book), args.Error(1)
}

func (m *MockBookRepository) FindByID(id string) (*entity.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *MockBookRepository) FindByTitle(title string) ([]*entity.Book, error) {
	args := m.Called(title)
	return args.Get(0).([]*entity.Book), args.Error(1)
}

func (m *MockBookRepository) FindByAuthor(author string) ([]*entity.Book, error) {
	args := m.Called(author)
	return args.Get(0).([]*entity.Book), args.Error(1)
}

func (m *MockBookRepository) FindByPublisher(publisher string) ([]*entity.Book, error) {
	args := m.Called(publisher)
	return args.Get(0).([]*entity.Book), args.Error(1)
}

func (m *MockBookRepository) Delete(id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func TestCreateBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewCreateBookUseCase(mockRepo)

	book := &entity.Book{
		ID:        uuid.New(),
		Title:     "Test Title",
		Author:    "Test Author",
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
	}

	mockRepo.On("Create", mock.Anything).Return(nil)

	_, err := u.CreateBook(usecases.CreateBookInput{
		Title:     book.Title,
		Author:    book.Author,
		Pages:     book.Pages,
		Publisher: book.Publisher,
		Year:      book.Year,
		ISBN:      book.ISBN,
	})

	if err != nil {
		t.Errorf("CreateBook() error = %v", err)
		return
	}

	mockRepo.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewDeleteBookUseCase(mockRepo)

	id := "test-id"

	mockRepo.On("Delete", id).Return(int64(1), nil)

	output, err := u.DeleteBook(usecases.DeleteBookInput{ID: id})

	if err != nil {
		t.Errorf("DeleteBook() error = %v", err)
		return
	}

	if output.RowsAffected != 1 {
		t.Errorf("DeleteBook() RowsAffected = %v, want 1", output.RowsAffected)
	}

	mockRepo.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewUpdateBookUseCase(mockRepo)

	id := uuid.New().String()
	input := usecases.UpdateBookInput{
		ID:        id,
		Title:     "Test Title",
		Author:    "Test Author",
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
	}

	mockRepo.On("Update", mock.Anything).Return(int64(1), nil)

	output, err := u.UpdateBook(input)

	if err != nil {
		t.Errorf("UpdateBook() error = %v", err)
		return
	}

	if output.RowsAffected != 1 {
		t.Errorf("UpdateBook() RowsAffected = %v, want 1", output.RowsAffected)
	}

	mockRepo.AssertExpectations(t)
}

func TestFindAllBooks(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewFindAllBooksUseCase(mockRepo)

	book := &entity.Book{
		ID:        uuid.New(),
		Title:     "Test Title",
		Author:    "Test Author",
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindAll").Return([]*entity.Book{book}, nil)

	output, err := u.FindAllBooks()

	if err != nil {
		t.Errorf("FindAllBooks() error = %v", err)
		return
	}

	if len(output.Books) != 1 {
		t.Errorf("FindAllBooks() len(Books) = %v, want 1", len(output.Books))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindBookByAuthor(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewFindBookByAuthorUseCase(mockRepo)

	author := "Test Author"
	book := &entity.Book{
		ID:        uuid.New(),
		Title:     "Test Title",
		Author:    author,
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByAuthor", author).Return([]*entity.Book{book}, nil)

	output, err := u.FindBookByAuthor(usecases.FindBookByAuthorInput{Author: author})

	if err != nil {
		t.Errorf("FindBookByAuthor() error = %v", err)
		return
	}

	if len(output.Books) != 1 {
		t.Errorf("FindBookByAuthor() len(Books) = %v, want 1", len(output.Books))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindBookByPublisher(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewFindBookByPublisherUseCase(mockRepo)

	publisher := "Test Publisher"
	book := &entity.Book{
		ID:        uuid.New(),
		Title:     "Test Title",
		Author:    "Test Author",
		Pages:     100,
		Publisher: publisher,
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByPublisher", publisher).Return([]*entity.Book{book}, nil)

	output, err := u.FindBookByPublisher(usecases.FindBookByPublisherInput{Publisher: publisher})

	if err != nil {
		t.Errorf("FindBookByPublisher() error = %v", err)
		return
	}

	if len(output.Books) != 1 {
		t.Errorf("FindBookByPublisher() len(Books) = %v, want 1", len(output.Books))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindBookByTitle(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewFindBookByTitleUseCase(mockRepo)

	title := "Test Title"
	book := &entity.Book{
		ID:        uuid.New(),
		Title:     title,
		Author:    "Test Author",
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByTitle", title).Return([]*entity.Book{book}, nil)

	output, err := u.FindBookByTitle(usecases.FindBookByTitleInput{Title: title})

	if err != nil {
		t.Errorf("FindBookByTitle() error = %v", err)
		return
	}

	if len(output.Books) != 1 {
		t.Errorf("FindBookByTitle() len(Books) = %v, want 1", len(output.Books))
	}

	mockRepo.AssertExpectations(t)
}

func TestFindBookByID(t *testing.T) {
	mockRepo := new(MockBookRepository)
	u := usecases.NewFindBookByIDUseCase(mockRepo)

	id := uuid.New().String()
	book := &entity.Book{
		ID:        uuid.MustParse(id),
		Title:     "Test Title",
		Author:    "Test Author",
		Pages:     100,
		Publisher: "Test Publisher",
		Year:      2022,
		ISBN:      "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByID", id).Return(book, nil)

	output, err := u.FindBookByID(usecases.FindBookByIDInput{ID: id})

	if err != nil {
		t.Errorf("FindBookByID() error = %v", err)
		return
	}

	if output.ID != id {
		t.Errorf("FindBookByID() ID = %v, want %v", output.ID, id)
	}

	mockRepo.AssertExpectations(t)
}
