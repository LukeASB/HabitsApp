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
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
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

func (ac *AuthController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	userRegisterRequest := data.RegisterUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&userRegisterRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	registeredUserData, err := ac.authModel.RegisterUserHandler(&userRegisterRequest)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.RegisterUserHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.RegisterUserHandler(registeredUserData)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.RegisterUserHandler - err: %s", err))
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

	ac.logger.DebugLog(fmt.Sprintf("authController.LoginHandler - Writing response: %s", response))
	numOfBytes, err := w.Write([]byte(response))
	ac.logger.DebugLog(fmt.Sprintf("authController.LoginHandler - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		ac.logger.ErrorLog(fmt.Sprintf("authController.LoginHandler - Error writing response: %s", err))
	}
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

	ac.logger.DebugLog(fmt.Sprintf("authController.LogoutHandler - Writing response: %s", response))
	numOfBytes, err := w.Write([]byte(response))
	ac.logger.DebugLog(fmt.Sprintf("authController.LogoutHandler - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		ac.logger.ErrorLog(fmt.Sprintf("authController.LogoutHandler - Error writing response: %s", err))
	}
}

// Called by frontend when the Short-Lived JWT expires and receives a 401 from the protected habits endpoints. Refreshes the authToken, refreshToken
func (ac *AuthController) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	userRefreshRequest := data.UserRefreshRequest{}

	if err := json.NewDecoder(r.Body).Decode(&userRefreshRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newAccessToken, err := ac.authModel.RefreshHandler(&userRefreshRequest, ac.jwtTokens)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.RefreshHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.RefreshHandler(&userRefreshRequest)

	if err != nil {
		ac.logger.DebugLog(fmt.Sprintf("authController.LogoutHandler - err: %s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", newAccessToken))
	ac.logger.DebugLog(fmt.Sprintf("authController.RefreshHandler - Writing response: %s", response))
	numOfBytes, err := w.Write([]byte(response))
	ac.logger.DebugLog(fmt.Sprintf("authController.RefreshHandler - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		ac.logger.ErrorLog(fmt.Sprintf("authController.RefreshHandler - Error writing response: %s", err))
	}
}
