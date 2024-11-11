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
	"time"
)

type AuthController struct {
	authModel  model.IAuthModel
	authView   view.IAuthView
	jwtTokens  session.IJWTTokens
	csrfTokens session.ICSRFToken
	logger     logger.ILogger
}

type IAuthController interface {
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

func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Validate the user - for new we'll have a test user - username: lsb, password: shushSecret
	//Create JWT authToken, refreshToken

	userAuth := data.UserAuth{}

	if err := json.NewDecoder(r.Body).Decode(&userAuth); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	accessToken, err := ac.jwtTokens.GenerateJSONWebTokens(userAuth.Username) // username will be passed as query param

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := ac.csrfTokens.CSRFToken(w); err != nil {
		http.Error(w, "Error generating CSRF token", http.StatusInternalServerError)
		return
	}

	response, err := ac.authView.LoginHandler(data.UserLoggedIn{Success: true, Username: userAuth.Username, LoggedInAt: time.Now()})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	w.Write(response)
}

func (ac *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// frontend will delete this short-lived JWT accessToken from session, delete user's refreshToken to prevent new short-lived JWT accessToken. It will invalidate itself after 5 mins.

	userLoggedOut := data.UserLoggedOut{}

	if err := json.NewDecoder(r.Body).Decode(&userLoggedOut); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ac.jwtTokens.DestroyJWTRefreshToken(userLoggedOut.Username)
	ac.csrfTokens.DestroyCSRFToken(w)

	w.Write([]byte("Logged out..."))
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
