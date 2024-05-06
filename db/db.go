package db

import (
	"database/sql"
	"fmt"
	"log"
	"yumandhika/golang-rest-api/types"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func NewMySQLStorage(config types.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func ConnectToSQLite(config types.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("sqlite", config.Name)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectToPostgreSQL(config types.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.User, config.Password, config.Name, config.Host, config.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
