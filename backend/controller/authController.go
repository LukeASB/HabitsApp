package controller

import (
	"dohabits/data"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
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

	accessToken, err := ac.jwtTokens.GenerateAccessJWT("testuser") // username will be passed as query param

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = ac.jwtTokens.GenerateRefreshJWT("testuser")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Generate CSRF Token
	csrfToken, err := ac.csrfTokens.GenerateCSRFToken()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Set the CSRF token as an HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		HttpOnly: true,
		Secure:   true, // If your site uses HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	response, err := ac.authView.LoginHandler(data.Login{Success: true, AccessToken: accessToken, LoggedInAt: time.Now()})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (ac *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// frontend will delete this short-lived JWT accessToken from session, delete user's refreshToken to prevent new short-lived JWT accessToken. It will invalidate itself after 5 mins.
}

func (ac *AuthController) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	// Called by frontend when the Short-Lived JWT expires and receives a 401 from the protected habits endpoints. Refreshes the authToken, refreshToken
	accessToken, err := ac.jwtTokens.RefreshJWTTokens()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Move to view layer
	response := map[string]interface{}{
		"success":     true,
		"accessToken": accessToken,
	}

	jsonRes, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(jsonRes)
}
