package controllers

import (
	"authentification/cmd/api/helpers"
	"authentification/data/models"
	"authentification/data/repository"
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	Client     *http.Client
	userModels repository.Repository
	json       *helpers.JsonResponse
}

func NewAuthenticationController(db *sql.DB) Authentication {
	return &authenticationController{
		userModels: models.New(db),
		json:       &helpers.JsonResponse{},
		Client:     &http.Client{},
	}
}

func (a *authenticationController) Authenticate(w http.ResponseWriter, r *http.Request) {
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

	valid, err := a.userModels.PasswordMatches(payload.Password, *user)
	if err != nil || !valid {
		_ = a.json.WriteJSONError(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	//log authentication
	err = a.logRequest("authentication", fmt.Sprintf("user %s authenticated", user.Email))
	if err != nil {
		log.Println("error logging authentication request: ", err)
	}

	response := &helpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}
	_ = response.WriteJSON(w, http.StatusOK, nil)
}

func (a *authenticationController) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/logger"

	req, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	_, err = a.Client.Do(req)

	return err

}
