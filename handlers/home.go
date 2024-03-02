package handlers

import (
    "net/http"

    "github.com/leggettc18/job-tracker/views"
)

func (h *Handler) HandleHome(writer http.ResponseWriter, request *http.Request) {
    views.Home().Render(request.Context(), writer)
}
