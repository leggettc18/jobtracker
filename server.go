package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
    addr string
    store Store
}

func NewServer(addr string, store Store) *Server {
    return &Server{addr: addr, store: store}
}

func (server *Server) Serve() {
    router := mux.NewRouter()

    //register services

    log.Printf("Starting server at %s\n", server.addr)
    log.Fatal(http.ListenAndServe(server.addr, router))
}
