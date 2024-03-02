package handlers

import (
    "log"
    "net/http"

    "github.com/leggettc18/job-tracker/views"
)

func (handler* Handler) HandleListJobs(writer http.ResponseWriter, request *http.Request) {
    isAddingJob := request.URL.Query().Get("isAddingJob") == "true"

    jobs, err := handler.store.GetJobs()
    if err != nil {
        log.Println(err)
        return
    }

    views.Jobs(jobs, isAddingJob).Render(request.Context(), writer)
}
