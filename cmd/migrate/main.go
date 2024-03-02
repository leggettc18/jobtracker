package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/leggettc18/job-tracker/store"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    connStr := "postgres://" + store.Envs.DBUser + ":" + store.Envs.DBPassword + "@" + store.Envs.DBAddress + "/" + store.Envs.DBName + "?sslmode=disable"
    storage := store.NewPostgresStorage(connStr)

    db, err := storage.Init()
    if err != nil {
        log.Fatal(err)
    }

    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatal(err)
    }
    m, err := migrate.NewWithDatabaseInstance(
        "file://cmd/migrate/migrations",
        "postgres",
        driver,
    )
    if err != nil {
        log.Fatal(err)
    }
    v, d, _ := m.Version()
    log.Printf("Version: %d, dirty: %v", v, d)

    cmd := os.Args[len(os.Args)-1]
    if cmd == "up" {
        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
            log.Fatal(err)
        }
    }
    if cmd == "down" {
        if err := m.Down(); err != nil && err != migrate.ErrNoChange {
            log.Fatal(err)
        }
    }
}
