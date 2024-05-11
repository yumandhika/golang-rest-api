package api

import (
	"database/sql"
	"log"
	"net/http"
	"yumandhika/golang-rest-api/services/users"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	userStore := users.NewStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	log.Println("Server Listening on", s.addr)
	return http.ListenAndServe(s.addr, v1)
}
