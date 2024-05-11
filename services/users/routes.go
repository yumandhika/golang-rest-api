package users

import (
	"log"
	"net/http"
	"yumandhika/golang-rest-api/utils"
)

type Handler struct {
	store UserStore
}

func NewHandler(store UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /users", h.handleGetUsers)
	router.HandleFunc("GET /users/{id}", h.handleGetUser)

}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("[ /users ]")
	utils.WriteJSON(w, http.StatusOK, "Success Get User")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("[ /users/%s ]", id)
	utils.WriteJSON(w, http.StatusOK, "Success Get Detail User")
}
