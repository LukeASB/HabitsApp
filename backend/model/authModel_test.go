package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/middleware/session"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
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
				Password:     "secretPassword012!",
				FirstName:    "First",
				LastName:     "Last",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			registeredUserData, err := authModel.RegisterUserHandler(val.userRegisterRequest)

			if err != nil {
				t.Errorf("TestRegisterUser Failed - err=%s", err)
				return
			}

			got := registeredUserData.Success

			if got != val.want {
				t.Errorf("TestRegisterUser Failed - got=%v, want=%v", got, val.want)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
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
			userAuth: &data.UserAuth{EmailAddress: "johndoe1@example.com", Password: "secretPassword012!"},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			userLoggedIn, err := authModel.LoginHandler(w, val.userAuth, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("TestLoginHandler Failed - err=%s", err)
				return
			}

			got := userLoggedIn.Success

			if got != val.want {
				t.Errorf("TestLoginHandler Failed - got=%v, want=%v", got, val.want)
			}
		})
	}
}

func TestLogoutHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

	csrfTokenMock := session.NewMockCSRFToken(logger)
	w := httptest.NewRecorder()

	testCases := []struct {
		name                 string
		userLoggedOutRequest *data.UserLoggedOutRequest
		want                 bool
	}{
		{
			name:                 "Successfully logout",
			userLoggedOutRequest: &data.UserLoggedOutRequest{UserID: "4", EmailAddress: "john.loggedin@example.com"},
			want:                 true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			userLoggedOutResponse, err := authModel.LogoutHandler(w, val.userLoggedOutRequest, jwtTokensMock, csrfTokenMock)

			if err != nil {
				t.Errorf("TestLogoutHandler Failed - err=%s", err)
				return
			}

			got := userLoggedOutResponse.Success

			if got != val.want {
				t.Errorf("TestLogoutHandler Failed - got=%v, want=%v", got, val.want)
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	authModel := NewAuthModel(logger, db)

	jwtTokensMock := session.NewMockJWTTokens("secretJwt")

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
			newAccessToken, err := authModel.RefreshHandler(val.userRefreshRequest, jwtTokensMock)

			if err != nil {
				t.Errorf("TestRefreshHandler Failed - err=%s", err)
				return
			}

			got := newAccessToken

			if val.want != got {
				t.Errorf("TestLogoutHandler Failed - got=%s, want=%s", got, val.want)
			}
		})
	}
}
