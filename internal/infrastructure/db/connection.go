package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gui-laranjeira/go-cleanarch/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := configs.GetDB()

	// TODO: handle connection string inside code

	// localhost connection string
	//connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Database)

	// container connection string
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.User, cfg.Pass, cfg.Container, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening connection with postgres: %v", err)
		return nil, err
	}

	err = db.Ping()
	return db, err
}
