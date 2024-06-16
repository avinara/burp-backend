package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/burp-backend/config"
)

type AuthControllerAPI interface {
	GoogleLogin(w http.ResponseWriter, r *http.Request)
	GoogleCallback(w http.ResponseWriter, r *http.Request)
}

type AuthController struct {
	config config.Config
}

func NewAuthController(config config.Config) *AuthController {
	return &AuthController{
		config: config,
	}
}

func (a *AuthController) GoogleLogin(w http.ResponseWriter, r *http.Request) {

	url := a.config.GoogleLoginConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a *AuthController) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != "randomstate" {
		return
	}

	code := r.URL.Query().Get("code")

	token, err := a.config.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(string(userData))
}
