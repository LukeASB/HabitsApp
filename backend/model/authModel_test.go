package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/helper"
	"dohabits/logger"
	"dohabits/middleware/session"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	testCases := []struct {
		name                string
		want                bool
		userRegisterRequest *data.RegisterUserRequest
	}{
		{
			name: "Successfully create a user",
			want: true,
			userRegisterRequest: &data.RegisterUserRequest{
				EmailAddress: "test@test.com",
				Password:     "1secret?Password",
				FirstName:    "First",
				LastName:     "Last",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			registeredUserData, err := authModel.RegisterUserHandler(val.userRegisterRequest)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			got := registeredUserData.Success

			if got != val.want {
				t.Errorf("%s - Failed - got=%v, want=%v", helper.GetFunctionName(), got, val.want)
				return
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name     string
		want     bool
		userAuth *data.UserAuth
	}{
		{
			name:     "Successfully login",
			want:     true,
			userAuth: &data.UserAuth{EmailAddress: "johndoe1@example.com", Password: "1secret?Password"},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			userLoggedIn, err := authModel.LoginHandler(w, val.userAuth, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			got := userLoggedIn.Success

			if got != val.want {
				t.Errorf("%s - Failed - got=%v, want=%v", helper.GetFunctionName(), got, val.want)
				return
			}
		})
	}
}

func TestLogoutHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name                 string
		userLoggedOutRequest *data.UserLoggedOutRequest
		want                 error
	}{
		{
			name:                 "Successfully logout",
			userLoggedOutRequest: &data.UserLoggedOutRequest{EmailAddress: "john.loggedin@example.com"},
			want:                 nil,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			var refreshTokenPath = "data/mock_refresh_tokens"
			var refreshTokenFile = "mock_refresh_token.txt"
			err := os.WriteFile(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, val.userLoggedOutRequest.EmailAddress, refreshTokenFile), []byte("testjwt"), 0644)
			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			err = authModel.LogoutHandler(w, val.userLoggedOutRequest, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name               string
		userRefreshRequest *data.UserRefreshRequest
		want               string
	}{
		{
			name:               "Get Refresh Token",
			userRefreshRequest: &data.UserRefreshRequest{EmailAddress: "johndoe1@example.com"},
			want:               "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es",
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			newAccessToken, _, err := authModel.RefreshHandler(w, val.userRefreshRequest, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			got := newAccessToken

			if val.want != got {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}
		})
	}
}
