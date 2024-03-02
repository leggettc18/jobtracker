package main

import (
    "log"

    "github.com/leggettc18/job-tracker/store"
)

func main () {
    connStr := "postgres://" + store.Envs.DBUser + ":" + store.Envs.DBPassword + "@" + store.Envs.DBAddress + "/" + store.Envs.DBName + "?sslmode=disable"
    storage := store.NewPostgresStorage(connStr)

    db, err := storage.Init()
    if err != nil {
        log.Fatal(err)
    }
    store_instance := store.NewStore(db)
    server := NewServer(":" + store.Envs.Port, store_instance)
    server.Serve()
}
