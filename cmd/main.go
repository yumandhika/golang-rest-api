package main

import (
	"database/sql"
	"log"
	"yumandhika/golang-rest-api/cmd/api"
	"yumandhika/golang-rest-api/configs"
	"yumandhika/golang-rest-api/db"
	"yumandhika/golang-rest-api/types"
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

	initStorage(db)
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
