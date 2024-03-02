package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


type PostgresStorage struct {
    db *sql.DB
}

func NewPostgresStorage(connString string) *PostgresStorage {
    db, err := sql.Open("postgres", connString)
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connected to Postgres!")
    return &PostgresStorage{db: db}
}

func (storage *PostgresStorage) Init() (*sql.DB, error) {
    // initialize tables

    return storage.db, nil
}
