package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/db"
	"github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository"
	usecases "github.com/gui-laranjeira/go-cleanarch/internal/usecases/stock"
)

func GetAllStockHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewStockSQLRepository(db)
	uc := usecases.NewFindAllStockUseCase(repo)

	output, err := uc.FindAllStock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting all stock entries: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetStockEntryByIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewStockSQLRepository(db)
	uc := usecases.NewFindStockByIDUseCase(repo)

	var input usecases.FindStockByIDInput
	input.StockID = chi.URLParam(r, "id")

	output, err := uc.FindStockByID(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting stock data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func GetStockEntryByBookIDHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewStockSQLRepository(db)
	uc := usecases.NewFindStockByBookIDUseCase(repo)

	var input usecases.FindStockByBookIDInput
	input.BookID = chi.URLParam(r, "id")

	output, err := uc.FindStockByBookID(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting stock data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func CreateStockEntryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewStockSQLRepository(db)
	uc := usecases.NewCreateStockEntryUseCase(repo)

	var input usecases.CreateStockEntryInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding request body: %v", err)
		return
	}

	output, err := uc.CreateStockEntry(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating stock entry: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func DeleteStockEntryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	repo := repository.NewStockSQLRepository(db)
	uc := usecases.NewDeleteStockEntryUseCase(repo)

	var input usecases.DeleteStockEntryInput
	input.StockID = chi.URLParam(r, "id")

	output, err := uc.DeleteStockEntry(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error deleting stock entry: %v", err)
		return
	}

	if output.RowsAffected <= 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

}

// TODO: Implement borrow and return book handlers
func BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	borrowRepo := repository.NewBorrowEntrySQLRepository(db)
	stockRepo := repository.NewStockSQLRepository(db)
	uc := usecases.NewBorrowBookUseCase(borrowRepo, stockRepo)

	var input usecases.BorrowBookInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding request body: %v", err)
		return
	}
	input.StockEntryID = chi.URLParam(r, "id")

	output, err := uc.BorrowBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error borrowing book: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.OpenConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while opening connection: %v", err)
		return
	}
	defer db.Close()

	borrowRepo := repository.NewBorrowEntrySQLRepository(db)
	stockRepo := repository.NewStockSQLRepository(db)
	uc := usecases.NewReturnBookUseCase(borrowRepo, stockRepo)

	var input usecases.ReturnBookInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error while decoding request body: %v", err)
		return
	}
	input.StockEntryID = chi.URLParam(r, "id")

	output, err := uc.ReturnBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error returning book: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
