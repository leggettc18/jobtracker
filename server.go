package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leggettc18/job-tracker/store"
    "github.com/leggettc18/job-tracker/handlers"
)

type Server struct {
    addr string
    store store.Store
}

func NewServer(addr string, store store.Store) *Server {
    return &Server{addr: addr, store: store}
}

func (server *Server) Serve() {
    router := mux.NewRouter()
    handler := handlers.New(&server.store)

    //register services
    router.HandleFunc("/", handler.HandleHome).Methods("GET")

    log.Printf("Starting server at %s\n", server.addr)
    log.Fatal(http.ListenAndServe(server.addr, router))
}
