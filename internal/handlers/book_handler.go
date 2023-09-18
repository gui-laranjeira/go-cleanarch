package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/book"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewBookSQLRepository(db)
	useCase := usecases.NewCreateBookUseCase(repo)

	var input usecases.CreateBookInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error while decoding book: %v", err)
		return
	}

	output, err := useCase.CreateBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while creating book: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewBookSQLRepository(db)
	useCase := usecases.NewFindAllBooksUseCase(repo)

	output, err := useCase.FindAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Printf("Error while getting all books: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewBookSQLRepository(db)
	useCase := usecases.NewFindBookByAuthorUseCase(repo)

	var input usecases.FindBookByAuthorInput

	input.Author = chi.URLParam(r, "author")

	if input.Author == "" || len(input.Author) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error while decoding book: %v", err)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error while decoding book: %v", err)
		return
	}

	output, err := useCase.FindBookByAuthor(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while getting book by author: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewBookSQLRepository(db)
	useCase := usecases.NewFindBookByIDUseCase(repo)

	idParam := chi.URLParam(r, "id")

	var input usecases.FindBookByIDInput
	input.ID = idParam

	output, err := useCase.FindBookByID(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while getting book by id: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}

	repo := repository.NewBookSQLRepository(db)
	usecase := usecases.NewDeleteBookUseCase(repo)

	idParam := chi.URLParam(r, "id")

	var input usecases.DeleteBookInput
	input.ID = idParam

	output, err := usecase.DeleteBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while deleting book: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

	if output.RowsAffected <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	repo := repository.NewBookSQLRepository(db)
	usecase := usecases.NewUpdateBookUseCase(repo)

	id := chi.URLParam(r, "id")

	oldBook, err := repo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Printf("Book not found: %v", err)
		return
	}

	var input usecases.UpdateBookInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	//#TODO Refactor this logic ->
	input.ID = id
	if input.Title == "" {
		input.Title = oldBook.Title
	}
	if input.Author == "" {
		input.Author = oldBook.Author
	}
	if input.Pages == 0 {
		input.Pages = oldBook.Pages
	}
	if input.Publisher == "" {
		input.Publisher = oldBook.Publisher
	}
	if input.Year == 0 {
		input.Year = oldBook.Year
	}
	if input.ISBN == "" {
		input.ISBN = oldBook.ISBN
	}

	output, err := usecase.UpdateBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while updating book: %v", err)
		return
	}

	if output.RowsAffected <= 0 {
		http.Error(w, "No rows affected", http.StatusNotFound)
		log.Print("No rows affected")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
