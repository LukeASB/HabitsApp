package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/helper"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/validation"
	"fmt"
	"net/http"
	"time"
)

type AuthModel struct {
	logger logger.ILogger
	db     db.IDB
}

type IAuthModel interface {
	RegisterUserHandler(userRegisterRequest *data.RegisterUserRequest) (*data.RegisterUserData, error)
	LoginHandler(w http.ResponseWriter, userAuth *data.UserAuth, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedInData, error)
	LogoutHandler(w http.ResponseWriter, UserLoggedOutRequest *data.UserLoggedOutRequest, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) error
	RefreshHandler(userRefreshRequest *data.UserRefreshRequest, jwtTokens session.IJWTTokens) (string, error)
}

func NewAuthModel(logger logger.ILogger, db db.IDB) *AuthModel {
	return &AuthModel{
		logger: logger,
		db:     db,
	}
}

func (am *AuthModel) RegisterUserHandler(userRegisterRequest *data.RegisterUserRequest) (*data.RegisterUserData, error) {
	am.logger.InfoLog(helper.GetFunctionName(), "")

	if !validation.IsValidName(userRegisterRequest.FirstName) {
		return nil, fmt.Errorf("%s - user: %s first name is invalid. FirstName: %s", helper.GetFunctionName(), userRegisterRequest.EmailAddress, userRegisterRequest.FirstName)
	}

	if !validation.IsValidName(userRegisterRequest.LastName) {
		return nil, fmt.Errorf("%s - user: %s last name is invalid. LastName: %s", helper.GetFunctionName(), userRegisterRequest.EmailAddress, userRegisterRequest.LastName)
	}

	if !validation.IsValidEmail(userRegisterRequest.EmailAddress) {
		return nil, fmt.Errorf("%s - user: %s email address is invalid. LastName: %s", helper.GetFunctionName(), userRegisterRequest.EmailAddress, userRegisterRequest.EmailAddress)
	}

	_, err := am.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userRegisterRequest.EmailAddress})

	if err != nil {
		return nil, err
	}

	hashedPassword, err := validation.HashPassword(userRegisterRequest.Password)

	if err != nil {
		return nil, err
	}

	userRegisterRequest.Password = string(hashedPassword)

	userData, err := am.db.RegisterUserHandler(userRegisterRequest)

	if err != nil {
		return nil, err
	}

	registerUserData, ok := userData.(*data.UserData)

	if !ok {
		return nil, fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	return &data.RegisterUserData{
		Success: true,
		User:    *registerUserData,
	}, nil
}

func (am *AuthModel) LoginHandler(w http.ResponseWriter, userAuth *data.UserAuth, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedInData, error) {
	am.logger.InfoLog(helper.GetFunctionName(), "")

	userDetails, err := am.db.GetUserDetails(userAuth)

	if err != nil {
		return nil, err
	}

	userData, ok := userDetails.(*data.UserData)

	if !ok {
		return nil, fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	if userData.IsLoggedIn {
		return nil, fmt.Errorf("%s - User is already logged in. UserID: %s", helper.GetFunctionName(), userData.UserID)
	}

	if !validation.VerifyUserPassword(userAuth.Password, userData.Password) {
		return nil, fmt.Errorf("%s - Invalid Password", helper.GetFunctionName())
	}

	accessToken, refreshToken, err := jwtTokens.GenerateJSONWebTokens(userAuth.EmailAddress)

	if err != nil {
		return nil, err
	}

	if err := csrfTokens.CSRFToken(w); err != nil {
		return nil, err
	}

	loggedInAt := time.Now()

	if err := am.db.LoginUser(&data.UserSession{
		UserID:       userData.UserID,
		RefreshToken: refreshToken,
		Device:       helper.GetDeviceInfo(),
		IPAddress:    helper.GetLocalIP(),
		CreatedAt:    loggedInAt,
	}); err != nil {
		return nil, err
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return &data.UserLoggedInData{
		Success:     true,
		User:        *userData,
		AccessToken: accessToken,
		LoggedInAt:  loggedInAt,
	}, nil
}

func (am *AuthModel) LogoutHandler(w http.ResponseWriter, userLoggedOutRequest *data.UserLoggedOutRequest, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) error {
	am.logger.InfoLog(helper.GetFunctionName(), "")

	userDetails, err := am.db.GetUserDetails(userLoggedOutRequest)

	if err != nil {
		return err
	}

	userData, ok := userDetails.(*data.UserData)

	if !ok {
		return fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	if !userData.IsLoggedIn {
		return fmt.Errorf("%s - User is not logged in. UserID: %s", helper.GetFunctionName(), userData.UserID)
	}

	jwtTokens.DestroyJWTRefreshToken(userLoggedOutRequest.EmailAddress)
	csrfTokens.DestroyCSRFToken(w)

	if err := am.db.LogoutUser(userData); err != nil {
		return err
	}

	return nil
}

func (am *AuthModel) RefreshHandler(userRefreshRequest *data.UserRefreshRequest, jwtTokens session.IJWTTokens) (string, error) {
	am.logger.InfoLog(helper.GetFunctionName(), "")

	newAccessToken, err := jwtTokens.RefreshJWTTokens(userRefreshRequest.EmailAddress)

	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
