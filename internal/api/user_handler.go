package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/kirildev25/go-api-pgx-postgres/internal/store"
	"github.com/kirildev25/go-api-pgx-postgres/internal/utils"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

func (h *UserHandler) validateRegisterUserRequest(request registerUserRequest) error {
	if request.Username == "" {
		return errors.New("Username is required")
	}

	if len(request.Username) > 50 {
		return errors.New("can't exceed 50 chars")
	}

	if request.Email == "" {
		return errors.New("Email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(request.Email) {
		return errors.New("Invalid email format")
	}

	if request.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("ERROR: decoding register request: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request payload"})
		return
	}

	err = h.validateRegisterUserRequest(req)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": err.Error()})
		return
	}

	user := store.User{
		Username: req.Username,
		Email:    req.Email,
	}

	if req.Bio != "" {
		user.Bio = req.Bio
	}
}
