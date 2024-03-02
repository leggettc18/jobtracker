package main

import (
	"github.com/gorilla/mux"
	"github.com/leggettc18/job-tracker/handlers"
)

type JobsService struct {
    handler *handlers.Handler
}

func NewJobsService(handler *handlers.Handler) *JobsService {
    return & JobsService{handler: handler}
}

func (service *JobsService) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/jobs", service.handler.HandleListJobs).Methods("GET")
    // r.HandleFunc("/jobs", service.handleCreateJob).Methods("POST")
    // r.HandleFunc("/jobs/{id}", service.handleGetJob).Methods("GET")
}
