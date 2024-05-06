package main

import (
	"database/sql"
	"log"
	"yumandhika/golang-rest-api/cmd/api"
	"yumandhika/golang-rest-api/db"
	"yumandhika/golang-rest-api/types"
)

func main() {

	sqliteConfig := types.DatabaseConfig{
		Name: "my_sqlite_db.db",
	}

	db, err := db.ConnectToSQLite(sqliteConfig)
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
