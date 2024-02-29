package main

import "log"

func main () {
    storage := NewPostgresStorage("postgres://jobtracker:jobtracker@localhost/jobtracker?sslmode=verify-full")

    db, err := storage.Init()
    if err != nil {
        log.Fatal(err)
    }
    store := NewStore(db)
    server := NewServer(":3000", store)
    server.Serve()
}
