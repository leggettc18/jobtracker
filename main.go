package main

import (
    "log"

    "github.com/leggettc18/job-tracker/store"
)

func main () {
    connStr := "postgres://" + Envs.DBUser + ":" + Envs.DBPassword + "@" + Envs.DBAddress + "/" + Envs.DBName + "?sslmode=disable"
    storage := store.NewPostgresStorage(connStr)

    db, err := storage.Init()
    if err != nil {
        log.Fatal(err)
    }
    store := store.NewStore(db)
    server := NewServer(":" + Envs.Port, store)
    server.Serve()
}
