package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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
