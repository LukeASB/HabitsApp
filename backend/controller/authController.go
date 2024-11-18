package controller

import (
	"dohabits/data"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthController struct {
	authModel  model.IAuthModel
	authView   view.IAuthView
	jwtTokens  session.IJWTTokens
	csrfTokens session.ICSRFToken
	logger     logger.ILogger
}

type IAuthController interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	RefreshHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthController(authModel model.IAuthModel, authView view.IAuthView, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken, logger logger.ILogger) *AuthController {
	return &AuthController{
		authModel:  authModel,
		authView:   authView,
		jwtTokens:  jwtTokens,
		csrfTokens: csrfTokens,
		logger:     logger,
	}
}

func (ac *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	userData := data.UserData{}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	registeredUserData, err := ac.authModel.RegisterUser(&userData)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.RegisterUser - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.RegisterUser(registeredUserData)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.RegisterUser - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userAuth := data.UserAuth{}

	if err := json.NewDecoder(r.Body).Decode(&userAuth); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userLoggedIn, err := ac.authModel.LoginHandler(w, &userAuth, ac.jwtTokens, ac.csrfTokens)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.LoginHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.LoginHandler(userLoggedIn)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.LoginHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", userLoggedIn.AccessToken))
	w.Write(response)
}

func (ac *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// frontend will delete this short-lived JWT accessToken from session, delete user's refreshToken to prevent new short-lived JWT accessToken. It will invalidate itself after 5 mins.

	userLoggedOutRequest := data.UserLoggedOutRequest{}

	if err := json.NewDecoder(r.Body).Decode(&userLoggedOutRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userLoggedOutResponse, err := ac.authModel.LogoutHandler(w, &userLoggedOutRequest, ac.jwtTokens, ac.csrfTokens)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.LogoutHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.LogoutHandler(userLoggedOutResponse)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.LogoutHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(response))
}

func (ac *AuthController) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	// Called by frontend when the Short-Lived JWT expires and receives a 401 from the protected habits endpoints. Refreshes the authToken, refreshToken
	newAccessToken, err := ac.jwtTokens.RefreshJWTTokens("username") // TO DO

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Move to view layer
	response := map[string]interface{}{
		"success": true,
	}

	jsonRes, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", newAccessToken))
	w.Write(jsonRes)
}
