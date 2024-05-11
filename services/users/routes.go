package users

import (
	"fmt"
	"log"
	"net/http"
	"yumandhika/golang-rest-api/services/auth"
	"yumandhika/golang-rest-api/types"
	"yumandhika/golang-rest-api/utils"

	"github.com/go-playground/validator"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
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

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
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
