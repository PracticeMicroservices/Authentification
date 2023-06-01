package controllers

import (
	"authentification/cmd/api/helpers"
	"authentification/data/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type Authentication interface {
	Authenticate(w http.ResponseWriter, r *http.Request)
}

type requestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authenticationController struct {
	userModels models.Models
	json       *helpers.JsonResponse
}

func NewAuthenticationController(db *sql.DB) Authentication {
	return &authenticationController{
		userModels: models.New(db),
		json:       &helpers.JsonResponse{},
	}
}

func (a *authenticationController) Authenticate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Authenticate")
	payload := &requestPayload{}
	err := helpers.ReadJSON(w, r, payload)
	if err != nil {
		_ = a.json.WriteJSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := a.userModels.GetByEmail(payload.Email)
	if err != nil {
		_ = a.json.WriteJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		_ = a.json.WriteJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	response := &helpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.User.Email),
	}
	_ = response.WriteJSON(w, http.StatusOK, nil)
}
