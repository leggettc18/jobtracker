package store

import "database/sql"

type Store interface {
    // Users
    CreateUser() error
}

type Storage struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
    return &Storage{db: db}
}

func (store *Storage) CreateUser() error {
    return nil
}
