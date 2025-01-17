package view

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"encoding/json"
)

type AuthView struct {
	logger logger.ILogger
}

type IAuthView interface {
	RegisterUserHandler(registeredUserData *data.RegisterUserData) ([]byte, error)
	LoginHandler(loginData *data.UserLoggedInData) ([]byte, error)
	RefreshHandler(userRefreshRequest *data.UserRefreshRequest) ([]byte, error)
}

func NewAuthView(logger logger.ILogger) *AuthView {
	return &AuthView{
		logger: logger,
	}
}

func (ac *AuthView) RegisterUserHandler(registeredUserData *data.RegisterUserData) ([]byte, error) {
	ac.logger.InfoLog(helper.GetFunctionName(), "")

	jsonRes, err := json.Marshal(data.RegisterUserResponse{
		Success: registeredUserData.Success,
		User: data.UserDataResponse{
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

func (ac *AuthView) LoginHandler(loginData *data.UserLoggedInData) ([]byte, error) {
	ac.logger.InfoLog(helper.GetFunctionName(), "")

	jsonRes, err := json.Marshal(data.UserLoggedInResponse{
		Success: loginData.Success,
		User: data.UserDataResponse{
			FirstName:    loginData.User.FirstName,
			LastName:     loginData.User.LastName,
			EmailAddress: loginData.User.EmailAddress,
			CreatedAt:    loginData.User.CreatedAt,
		},
		LoggedInAt: loginData.LoggedInAt,
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, nil
}

func (ac *AuthView) RefreshHandler(userRefreshRequest *data.UserRefreshRequest) ([]byte, error) {
	ac.logger.InfoLog(helper.GetFunctionName(), "")

	jsonRes, err := json.Marshal(data.UserRefreshResponse{
		Success:      true,
		EmailAddress: userRefreshRequest.EmailAddress,
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, err
}
