package view

import (
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"time"
)

type AuthView struct {
	logger logger.ILogger
}

type IAuthView interface {
	RegisterUserHandler(registeredUserData *data.RegisterUserResponse) ([]byte, error)
	LoginHandler(loginData *data.UserLoggedIn) ([]byte, error)
	LogoutHandler(logoutData *data.UserLoggedOutResponse) ([]byte, error)
	RefreshHandler(userRefreshRequest *data.UserRefreshRequest, accessToken string) ([]byte, error)
}

func NewAuthView(logger logger.ILogger) *AuthView {
	return &AuthView{
		logger: logger,
	}
}

func (ac *AuthView) RegisterUserHandler(registeredUserData *data.RegisterUserResponse) ([]byte, error) {
	ac.logger.InfoLog("authView.RegisterUserHandler")

	jsonRes, err := json.Marshal(struct {
		Success bool `json:"Succcess"`
		User    struct {
			UserID       string    `json:"UserID"`
			FirstName    string    `json:"FirstName"`
			LastName     string    `json:"LastName"`
			EmailAddress string    `json:"EmailAddress"`
			CreatedAt    time.Time `json:"CreatedAt"`
		} `json:"User"`
	}{
		Success: registeredUserData.Success,
		User: struct {
			UserID       string    `json:"UserID"`
			FirstName    string    `json:"FirstName"`
			LastName     string    `json:"LastName"`
			EmailAddress string    `json:"EmailAddress"`
			CreatedAt    time.Time `json:"CreatedAt"`
		}{
			UserID:       registeredUserData.User.UserID,
			FirstName:    registeredUserData.User.FirstName,
			LastName:     registeredUserData.User.LastName,
			EmailAddress: registeredUserData.User.EmailAddress,
			CreatedAt:    registeredUserData.User.CreatedAt,
		},
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, err
}

func (ac *AuthView) LoginHandler(loginData *data.UserLoggedIn) ([]byte, error) {
	ac.logger.InfoLog("authView.LoginHandler")

	jsonRes, err := json.Marshal(struct {
		Success bool `json:"Succcess"`
		User    struct {
			UserID       string    `json:"UserID"`
			FirstName    string    `json:"FirstName"`
			LastName     string    `json:"LastName"`
			EmailAddress string    `json:"EmailAddress"`
			CreatedAt    time.Time `json:"CreatedAt"`
		} `json:"User"`
		AccessToken string    `json:"AccessToken"`
		LoggedInAt  time.Time `json:"LoggedInAt"`
	}{
		Success: loginData.Success,
		User: struct {
			UserID       string    `json:"UserID"`
			FirstName    string    `json:"FirstName"`
			LastName     string    `json:"LastName"`
			EmailAddress string    `json:"EmailAddress"`
			CreatedAt    time.Time `json:"CreatedAt"`
		}{
			UserID:       loginData.User.UserID,
			FirstName:    loginData.User.FirstName,
			LastName:     loginData.User.LastName,
			EmailAddress: loginData.User.EmailAddress,
			CreatedAt:    loginData.User.CreatedAt,
		},
		AccessToken: loginData.AccessToken,
		LoggedInAt:  loginData.LoggedInAt,
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, nil
}

func (ac *AuthView) LogoutHandler(logoutData *data.UserLoggedOutResponse) ([]byte, error) {
	ac.logger.InfoLog("authView.LogoutHandler")

	jsonRes, err := json.Marshal(logoutData)

	if err != nil {
		return nil, err
	}

	return jsonRes, err
}

func (ac *AuthView) RefreshHandler(userRefreshRequest *data.UserRefreshRequest, accessToken string) ([]byte, error) {
	ac.logger.InfoLog("authView.RefreshHandler")

	jsonRes, err := json.Marshal(data.UserRefreshResponse{
		Success:      true,
		EmailAddress: userRefreshRequest.EmailAddress,
		AccessToken:  accessToken,
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, err
}
