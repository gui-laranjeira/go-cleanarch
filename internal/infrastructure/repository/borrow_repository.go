package repository

import (
	"database/sql"
	"time"

	"github.com/gui-laranjeira/go-cleanarch/internal/entity"
	repository "github.com/gui-laranjeira/go-cleanarch/internal/infrastructure/repository/errors"
)

type BorrowSQLRepository struct {
	db *sql.DB
}

func NewBorrowEntrySQLRepository(db *sql.DB) entity.IBorrowEntryRepository {
	return &BorrowSQLRepository{
		db: db,
	}
}

func (r *BorrowSQLRepository) Borrow(entry *entity.BorrowEntry) (int64, error) {
	query := `INSERT INTO borrow_history (id_costumer, id_stock_entry, borrow_date, due_date, returned, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	rows, err := r.db.Exec(query, entry.CostumerID, entry.StockEntryID, entry.BorrowDate, entry.DueDate, entry.Returned, entry.CreatedAt, entry.UpdatedAt)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return rowsAffected, repository.NoRowsAffectedError()
	}

	if rowsAffected > 1 {
		return rowsAffected, repository.MultipleRowsAffectedError()
	}

	return rowsAffected, nil
}

func (r *BorrowSQLRepository) Return(entry *entity.BorrowEntry) (int64, error) {
	query := `UPDATE borrow_history SET returned = true, updated_at = $1 WHERE id_costumer = $2 AND id_stock_entry = $3 and returned = false`
	rows, err := r.db.Exec(query, time.Now(), entry.CostumerID, entry.StockEntryID)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return rowsAffected, repository.NoRowsAffectedError()
	}

	if rowsAffected > 1 {
		return rowsAffected, repository.MultipleRowsAffectedError()
	}

	return rowsAffected, nil
}

func (r *BorrowSQLRepository) GetBorrowEntryByID(id_stock_entry string, id_costumer string) (*entity.BorrowEntry, error) {
	query := `SELECT * FROM borrow_history WHERE id_stock_entry = $1 AND id_costumer = $2`
	rows, err := r.db.Query(query, id_stock_entry, id_costumer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowEntry entity.BorrowEntry
	for rows.Next() {
		err = rows.Scan(&borrowEntry.CostumerID, &borrowEntry.StockEntryID, &borrowEntry.BorrowDate, &borrowEntry.DueDate, &borrowEntry.Returned, &borrowEntry.CreatedAt, &borrowEntry.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &borrowEntry, nil
}
