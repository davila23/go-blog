package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/davila23/jwt-auth/api/auth"

	"github.com/davila23/jwt-auth/api/models"
	"github.com/davila23/jwt-auth/api/responses"
)

// Login is the signIn method
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}
