package store

import (
	"database/sql"

	"github.com/leggettc18/job-tracker/types"
)

type Store interface {
    // Jobs
    GetJobs() ([]types.Job, error)
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
