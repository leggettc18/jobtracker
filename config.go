package main

import (
	"fmt"
	"os"
)

type Config struct {
    Port string
    DBUser string
    DBPassword string
    DBAddress string
    DBName string
    JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
    return Config{
        Port: getEnv("JOBTRACKER_PORT", "8080"),
        DBUser: getEnv("JOBTRACKER_DB_USER", "jobtracker"),
        DBPassword: getEnv("JOBTRACKER_DB_PASSWORD", "jobtracker"),
        DBAddress: fmt.Sprintf("%s:%s", getEnv("JOBTRACKER_DB_HOST", "127.0.0.1"), getEnv("JOBTRACKER_DB_PORT", "5432")),
        DBName: getEnv("JOBTRACKER_DB_NAME", "jobtracker"),
        JWTSecret: getEnv("JOBTRACKER_SECRET", "randomsecretkey"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    } 
    return fallback
}
