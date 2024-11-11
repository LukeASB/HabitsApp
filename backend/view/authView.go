package view

import (
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"net/http"
)

type AuthView struct {
	logger logger.ILogger
}

type IAuthView interface {
	LoginHandler(loginData data.UserLoggedIn) ([]byte, error)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	RefreshHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthView(logger logger.ILogger) *AuthView {
	return &AuthView{
		logger: logger,
	}
}

func (ac *AuthView) LoginHandler(loginData data.UserLoggedIn) ([]byte, error) {
	ac.logger.InfoLog("authView.LoginHandler")

	jsonRes, err := json.Marshal(loginData)

	if err != nil {
		return nil, err
	}

	return jsonRes, nil
}

func (ac *AuthView) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ac.logger.InfoLog("authView.LogoutHandler")
}

func (ac *AuthView) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	ac.logger.InfoLog("authView.RefreshHandler")
}
