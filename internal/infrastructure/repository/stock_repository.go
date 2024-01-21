package repository

import (
	"database/sql"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	repository "github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository/errors"
)

type StockSQLRepository struct {
	db *sql.DB
}

func NewStockSQLRepository(db *sql.DB) entity.IStockRepository {
	return &StockSQLRepository{
		db: db,
	}
}

func (r *StockSQLRepository) CreateStockEntry(entry *entity.StockEntry) error {
	query := `INSERT INTO stock (id_stock_entry, id_book, available) VALUES ($1, $2, $3)`
	rows, err := r.db.Exec(query, entry.StockID, entry.BookID, entry.Available)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.NoRowsAffectedError()
	}

	if rowsAffected > 1 {
		return repository.MultipleRowsAffectedError()
	}

	return nil
}

func (r *StockSQLRepository) GetAllStock() ([]*entity.StockEntry, error) {
	query := `SELECT * FROM stock`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stock []*entity.StockEntry
	for rows.Next() {
		var stockEntry entity.StockEntry
		err = rows.Scan(&stockEntry.StockID, &stockEntry.BookID, &stockEntry.Available)
		if err != nil {
			return nil, err
		}
		stock = append(stock, &stockEntry)
	}

	return stock, nil
}

func (r *StockSQLRepository) GetStockEntryByID(id string) (*entity.StockEntry, error) {
	query := `SELECT * FROM stock WHERE id_stock_entry = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var stockEntry entity.StockEntry
	for rows.Next() {
		err = rows.Scan(&stockEntry.StockID, &stockEntry.BookID, &stockEntry.Available)
		if err != nil {
			return nil, err
		}
	}

	return &stockEntry, nil
}

func (r *StockSQLRepository) GetStockEntryByBookID(id_book string) ([]*entity.StockEntry, error) {
	query := `SELECT * FROM stock WHERE id_book = $1`
	rows, err := r.db.Query(query, id_book)
	if err != nil {
		return nil, err
	}

	var stock []*entity.StockEntry
	for rows.Next() {
		var stockEntry entity.StockEntry
		err = rows.Scan(&stockEntry.StockID, &stockEntry.BookID, &stockEntry.Available)
		if err != nil {
			return nil, err
		}
		stock = append(stock, &stockEntry)
	}

	return stock, nil
}

func (r *StockSQLRepository) DeleteStockEntry(id_stock_entry string) (int64, error) {
	query := `DELETE FROM stock WHERE id_stock_entry = $1`
	rows, err := r.db.Exec(query, id_stock_entry)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *StockSQLRepository) BorrowBook(id_stock_entry string) error {
	query := `UPDATE stock SET available = false WHERE id_stock_entry = $1`
	_, err := r.db.Exec(query, id_stock_entry)
	if err != nil {
		return err
	}
	return nil
}
func (r *StockSQLRepository) ReturnBook(id_stock_entry string) error {
	query := `UPDATE stock SET available = true WHERE id_stock_entry = $1`
	_, err := r.db.Exec(query, id_stock_entry)
	if err != nil {
		return err
	}

	return nil
}
