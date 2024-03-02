package handlers

import (
    "github.com/leggettc18/job-tracker/store"
)

type Handler struct {
    store *store.Store
}

func New(store *store.Store) *Handler {
    return &Handler{
        store: store,
    }
}
