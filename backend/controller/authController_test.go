package controller

import (
	"bytes"
	"context"
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name                string
		want                bool
		userRegisterRequest *data.RegisterUserRequest
	}{
		{
			name: "Successfully create a user",
			want: true,
			userRegisterRequest: &data.RegisterUserRequest{
				EmailAddress: "controlller@test.com",
				Password:     "1secret?Password",
				FirstName:    "First",
				LastName:     "Last",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledUserRegisterRequest, err := json.Marshal(val.userRegisterRequest)

			if err != nil {
				t.Errorf("TestRegisterUserHandler err: %s", err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/register", endpoint), io.NopCloser(bytes.NewBuffer(marshalledUserRegisterRequest)))
			w := httptest.NewRecorder()

			authController.RegisterUserHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestRegisterUserHandler - HTTP Status Code = %d", status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("TestRegisterUserHandler err: %s", err)
				return
			}

			registerUserResponseData := &data.RegisterUserResponse{}

			err = json.Unmarshal(got, registerUserResponseData)

			if err != nil {
				t.Errorf("TestRegisterUserHandler err: %s", err)
				return
			}

			if registerUserResponseData.User.FirstName != val.userRegisterRequest.FirstName {
				t.Errorf("TestRegisterUserHandler - first name does not match registration. got=%s want=%s", registerUserResponseData.User.FirstName, val.userRegisterRequest.FirstName)
				return
			}

			if registerUserResponseData.User.LastName != val.userRegisterRequest.LastName {
				t.Errorf("TestRegisterUserHandler - last name does not match registration. got=%s want=%s", registerUserResponseData.User.LastName, val.userRegisterRequest.LastName)
				return
			}

			if registerUserResponseData.User.EmailAddress != val.userRegisterRequest.EmailAddress {
				t.Errorf("TestRegisterUserHandler - email does not match registration. got=%s want=%s", registerUserResponseData.User.EmailAddress, val.userRegisterRequest.EmailAddress)
				return
			}

			if registerUserResponseData.Success != val.want {
				t.Errorf("TestRegisterUser Failed - got=%v, want=%v", registerUserResponseData.Success, val.want)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name     string
		want     bool
		userAuth *data.UserAuth
	}{
		{
			name: "Successfully login",
			want: true,
			userAuth: &data.UserAuth{
				EmailAddress: "johndoe1@example.com",
				Password:     "1secret?Password",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledUserAuth, err := json.Marshal(val.userAuth)

			if err != nil {
				t.Errorf("TestLoginHandler err: %s", err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/login", endpoint), io.NopCloser(bytes.NewBuffer(marshalledUserAuth)))
			w := httptest.NewRecorder()

			authController.LoginHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestLoginHandler - HTTP Status Code = %d", status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("TestLoginHandler err: %s", err)
				return
			}

			userLoggedInResponse := &data.UserLoggedInResponse{}

			err = json.Unmarshal(got, userLoggedInResponse)

			if err != nil {
				t.Errorf("TestLoginHandler err: %s", err)
				return
			}

			if userLoggedInResponse.User.EmailAddress != val.userAuth.EmailAddress {
				t.Errorf("TestLoginHandler - email does not match logged in user. got=%s want=%s", userLoggedInResponse.User.EmailAddress, val.userAuth.EmailAddress)
				return
			}

			if userLoggedInResponse.Success != val.want {
				t.Errorf("TestRegisterUser Failed - got=%v, want=%v", userLoggedInResponse.Success, val.want)
			}

		})
	}
}

func TestLogoutHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name     string
		want     string
		username string
	}{
		{
			name:     "Successfully logout",
			want:     "200 OK",
			username: "john.loggedin@example.com",
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			var refreshTokenPath = "data/mock_refresh_tokens"
			var refreshTokenFile = "mock_refresh_token.txt"
			err := os.WriteFile(fmt.Sprintf("../%s/%s_%s", refreshTokenPath, val.username, refreshTokenFile), []byte("testjwt"), 0644)
			if err != nil {
				t.Errorf("TestLogoutHandler Failed - err=%s", err)
			}

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/logout", endpoint), nil)
			claims := &session.Claims{Username: val.username}
			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)
			req = req.WithContext(ctx)
			w := httptest.NewRecorder()

			authController.LogoutHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestLogoutHandler - HTTP Status Code = %d", status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got := res.Status

			if val.want != got {
				t.Errorf("TestLogoutHandler Failed - got=%s, want=%s", got, val.want)
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name               string
		want               bool
		userRefreshRequest *data.UserRefreshRequest
	}{
		{
			name: "Successfully Get a Refresh Token",
			want: true,
			userRefreshRequest: &data.UserRefreshRequest{
				EmailAddress: "johndoe1@example.com",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledUserRefreshRequest, err := json.Marshal(val.userRefreshRequest)

			if err != nil {
				t.Errorf("TestRefreshHandler err: %s", err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/logout", endpoint), io.NopCloser(bytes.NewBuffer(marshalledUserRefreshRequest)))
			w := httptest.NewRecorder()

			authController.RefreshHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestRefreshHandler - HTTP Status Code = %d", status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("TestRefreshHandler err: %s", err)
				return
			}

			userRefreshResponse := &data.UserRefreshResponse{}

			err = json.Unmarshal(got, userRefreshResponse)

			if userRefreshResponse.EmailAddress != val.userRefreshRequest.EmailAddress {
				t.Errorf("TestRefreshHandler - email does not match logged out  user. got=%s want=%s", userRefreshResponse.EmailAddress, val.userRefreshRequest.EmailAddress)
				return
			}

			if userRefreshResponse.Success != val.want {
				t.Errorf("TestRefreshHandler Failed")
			}
		})
	}
}
