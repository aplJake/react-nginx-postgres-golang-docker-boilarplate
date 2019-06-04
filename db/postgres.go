package db

import (
		"database/sql"
		_ "github.com/lib/pq"
)

type PostgresRepository struct {
		db *sql.DB
}

func NewPostgresDB(url string) (*PostgresRepository, error)  {
		db, err := sql.Open("postgres", url)
		if err != nil {
				return nil, err
		}
		err = db.Ping()
		if err != nil {
				return nil, err
		}
		return &PostgresRepository{db}, nil
}

func (db *PostgresRepository) GetDB() *sql.DB  {
		return db.db
}

func (db *PostgresRepository) Close()  {
		db.db.Close()
}