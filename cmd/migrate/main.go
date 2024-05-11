package main

import (
	"log"
	"os"
	"yumandhika/golang-rest-api/configs"
	"yumandhika/golang-rest-api/db"
	"yumandhika/golang-rest-api/types"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg := types.DatabaseConfig{
		User:     configs.Envs.DBUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBHost,
		Port:     int(configs.Envs.DBPort),
		Name:     configs.Envs.DBName,
	}

	db, err := db.ConnectToPostgreSQL(cfg)
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
