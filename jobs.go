package main

import (
	"github.com/gorilla/mux"
	"github.com/leggettc18/job-tracker/store"
)

type JobsService struct {
    store store.Store
}

func NewJobsService(store store.Store) *JobsService {
    return & JobsService{store: store}
}

func (service *JobsService) RegisterRoutes(router *mux.Router) {
    // r.HandleFunc("/jobs", service.handleCreateJob).Methods("POST")
    // r.HandleFunc("/jobs/{id}", service.handleGetJob).Methods("GET")
}
