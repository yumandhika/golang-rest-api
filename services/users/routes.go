package users

import (
	"log"
	"net/http"
	"yumandhika/golang-rest-api/types"
	"yumandhika/golang-rest-api/utils"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /users", h.handleGetUsers)
	router.HandleFunc("GET /users/{id}", h.handleGetUser)
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("[ /login ]")
	utils.WriteJSON(w, http.StatusOK, "Success Login")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("[ /register ]")
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Success Register")
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
