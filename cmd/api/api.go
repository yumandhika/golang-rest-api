package api

import (
	"log"
	"net/http"
	"yumandhika/golang-rest-api/services/users"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userHandler := users.NewHandler()
	userHandler.RegisterRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	log.Println("Server Listening on", s.addr)
	return http.ListenAndServe(s.addr, v1)
}
