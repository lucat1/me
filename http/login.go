package http

import (
	"encoding/json"
	"net/http"
)

type LoginFormData struct {
	Username      string
	UsernameError string
	Password      string
	PasswordError string
	Error         string
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r)
	logger.Info("Rendering login index")

	if err := RenderPage[LoginFormData](w, r, "login", "Login", LoginFormData{}); err != nil {
		logger.With("err", err).Error("Could not render page")
	}
}

type LoginDoData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	logger := GetLogger(r)
	var loginData LoginDoData

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		// TODO: display errors
		logger.With("err", err).Warn("Invalid login data")
		return
	}

	logger.With("username", loginData.Username).Info("Login attempt")

	// Basic data to persist form values
	data := LoginFormData{
		Username: loginData.Username,
		Password: loginData.Password,
	}

	// TODO: call the login script instead of aribtrarely setting the error
	data.Error = "Not implemented yet"
	partial, err := RenderBlock[LoginFormData]("login.form", data)
	if err != nil {
		logger.With("err", err).Error("Could not render page")
	}

	if _, err = w.Write(partial); err != nil {
		logger.With("err", err).Error("Could not render login partial")
	}
}
