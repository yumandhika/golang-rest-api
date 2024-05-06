package main

import (
	"log"
	"os"
	"yumandhika/golang-rest-api/configs"
	"yumandhika/golang-rest-api/db"
	"yumandhika/golang-rest-api/types"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := types.DatabaseConfig{
		User:     configs.Envs.DBUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBAddress,
		Port:     configs.Envs.DBPort,
		Name:     configs.Envs.DBName,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
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
