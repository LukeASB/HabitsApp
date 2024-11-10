package model

import (
	"dohabits/db"
	"dohabits/logger"
	"net/http"
)

type AuthModel struct {
	logger logger.ILogger
	db     db.IDB
}

type IAuthModel interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	RefreshHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthModel(logger logger.ILogger, db db.IDB) *AuthModel {
	return &AuthModel{
		logger: logger,
		db:     db,
	}
}

func (ac *AuthModel) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ac.logger.InfoLog("authModel.LoginHandler")
}

func (ac *AuthModel) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ac.logger.InfoLog("authModel.LogoutHandler")
}

func (ac *AuthModel) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	ac.logger.InfoLog("authModel.RefreshHandler")
}
