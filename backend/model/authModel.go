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
	RegisterUser(userData *data.UserData) (*data.RegisterUser, error)
	LoginHandler(w http.ResponseWriter, userAuth *data.UserAuth, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedIn, error)
	LogoutHandler(w http.ResponseWriter, UserLoggedOutRequest *data.UserLoggedOutRequest, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedOutResponse, error)
	RefreshHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthModel(logger logger.ILogger, db db.IDB) *AuthModel {
	return &AuthModel{
		logger: logger,
		db:     db,
	}
}

func (am *AuthModel) RegisterUser(userData *data.UserData) (*data.RegisterUser, error) {
	am.logger.InfoLog("authModel.RegisterUser")

	if !validation.IsValidName(userData.FirstName) {
		return nil, fmt.Errorf("authModel.RegisterUser - user: %s first name is invalid. FirstName: %s", userData.UserID, userData.FirstName)
	}

	if !validation.IsValidName(userData.LastName) {
		return nil, fmt.Errorf("authModel.RegisterUser - user: %s last name is invalid. LastName: %s", userData.UserID, userData.LastName)
	}

	if !validation.IsValidEmail(userData.EmailAddress) {
		return nil, fmt.Errorf("authModel.RegisterUser - user: %s email address is invalid. LastName: %s", userData.UserID, userData.EmailAddress)
	}

	userDetails, err := am.db.GetUserDetails(&data.UserAuth{EmailAddress: userData.EmailAddress})

	if err != nil {
		return nil, err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return nil, fmt.Errorf("authModel.RegisterUser - data.UserData is invalid")
	}

	if len(currentUserData.UserID) > 0 {
		return nil, fmt.Errorf("authModel.RegisterUser - User already exists. UserID: %s, EmailAddress: %s", currentUserData.UserID, currentUserData.EmailAddress)
	}

	hashedPassword, err := validation.HashPassword(userData.Password)

	if err != nil {
		return nil, err
	}

	userData.Password = string(hashedPassword)

	if err := am.db.RegisterUser(userData); err != nil {
		return nil, err
	}

	return &data.RegisterUser{
		Success: true,
		User:    *userData,
	}, nil
}

func (am *AuthModel) LoginHandler(w http.ResponseWriter, userAuth *data.UserAuth, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedIn, error) {
	am.logger.InfoLog("authModel.LoginHandler")

	userDetails, err := am.db.GetUserDetails(userAuth)

	if err != nil {
		return nil, err
	}

	userData, ok := userDetails.(data.UserData)

	if !ok {
		return nil, fmt.Errorf("authModel.LoginHandler - data.UserData is invalid")
	}

	if userData.IsLoggedIn {
		return nil, fmt.Errorf("authModel.LoginHandler - User is already logged in. UserID: %s", userData.UserID)
	}

	if !validation.VerifyUserPassword(userAuth.Password, userData.Password) {
		return nil, fmt.Errorf("authModel.LoginHandler - Invalid Password")
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

	return &data.UserLoggedIn{
		Success:     true,
		User:        userData,
		AccessToken: accessToken,
		LoggedInAt:  loggedInAt,
	}, nil
}

func (am *AuthModel) LogoutHandler(w http.ResponseWriter, UserLoggedOutRequest *data.UserLoggedOutRequest, jwtTokens session.IJWTTokens, csrfTokens session.ICSRFToken) (*data.UserLoggedOutResponse, error) {
	am.logger.InfoLog("authModel.LogoutHandler")

	userDetails, err := am.db.GetUserDetails(UserLoggedOutRequest)

	if err != nil {
		return nil, err
	}

	userData, ok := userDetails.(data.UserData)

	if !ok {
		return nil, fmt.Errorf("authModel.LoginHandler - data.UserData is invalid")
	}

	if !userData.IsLoggedIn {
		return nil, fmt.Errorf("authModel.LoginHandler - User is not logged in. UserID: %s", userData.UserID)
	}

	jwtTokens.DestroyJWTRefreshToken(UserLoggedOutRequest.EmailAddress)
	csrfTokens.DestroyCSRFToken(w)

	if err := am.db.LogoutUser(UserLoggedOutRequest); err != nil {
		return nil, err
	}

	return &data.UserLoggedOutResponse{
		Success:      true,
		UserID:       UserLoggedOutRequest.UserID,
		EmailAddress: UserLoggedOutRequest.EmailAddress,
		LoggedOutAt:  time.Now(),
	}, nil
}

func (am *AuthModel) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	am.logger.InfoLog("authModel.RefreshHandler")
	// if an API call fails returns forbidden/unauth, frontend calls this to re-issue a short-lived token. May not need as middleware should handle if subscribed.
}
