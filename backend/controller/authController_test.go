package controller

import (
	"bytes"
	"context"
	"dohabits/data"
	"dohabits/db"
	"dohabits/helper"
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
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, "/register", io.NopCloser(bytes.NewBuffer(marshalledUserRegisterRequest)))
			w := httptest.NewRecorder()

			authController.RegisterUserHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			registerUserResponseData := &data.RegisterUserResponse{}

			err = json.Unmarshal(got, registerUserResponseData)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if registerUserResponseData.User.FirstName != val.userRegisterRequest.FirstName {
				t.Errorf("%s - Failed - first name does not match registration - got=%s, want=%s", helper.GetFunctionName(), registerUserResponseData.User.FirstName, val.userRegisterRequest.FirstName)
				return
			}

			if registerUserResponseData.User.LastName != val.userRegisterRequest.LastName {
				t.Errorf("%s - Failed - last name does not match registration - got=%s, want=%s", helper.GetFunctionName(), registerUserResponseData.User.LastName, val.userRegisterRequest.LastName)
				return
			}

			if registerUserResponseData.User.EmailAddress != val.userRegisterRequest.EmailAddress {
				t.Errorf("%s - Failed - email does not match registration - got=%s, want=%s", helper.GetFunctionName(), registerUserResponseData.User.EmailAddress, val.userRegisterRequest.EmailAddress)
				return
			}

			if registerUserResponseData.Success != val.want {
				t.Errorf("%s - Failed - got=%v, want=%v", helper.GetFunctionName(), registerUserResponseData.Success, val.want)
				return
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, "/login", io.NopCloser(bytes.NewBuffer(marshalledUserAuth)))
			w := httptest.NewRecorder()

			authController.LoginHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			userLoggedInResponse := &data.UserLoggedInResponse{}

			err = json.Unmarshal(got, userLoggedInResponse)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if userLoggedInResponse.User.EmailAddress != val.userAuth.EmailAddress {
				t.Errorf("%s - Failed - email does not match logged in user - got=%s, want=%s", helper.GetFunctionName(), userLoggedInResponse.User.EmailAddress, val.userAuth.EmailAddress)
				return
			}

			if userLoggedInResponse.Success != val.want {
				t.Errorf("%s - Failed - got=%v, want=%v", helper.GetFunctionName(), userLoggedInResponse.Success, val.want)
				return
			}

		})
	}
}

func TestLogoutHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
			}

			req := httptest.NewRequest(http.MethodPost, "/logout", nil)
			claims := &session.Claims{Username: val.username}
			ctx := context.WithValue(req.Context(), session.ClaimsKey, claims)
			req = req.WithContext(ctx)
			w := httptest.NewRecorder()

			authController.LogoutHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got := res.Status

			if val.want != got {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), got, val.want)
				return
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := logger.NewLogger(0)
	db := db.NewMockDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

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
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, "/refresh", io.NopCloser(bytes.NewBuffer(marshalledUserRefreshRequest)))
			w := httptest.NewRecorder()

			authController.RefreshHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("%s - Failed - HTTP Status Code = %d", helper.GetFunctionName(), status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			userRefreshResponse := data.UserRefreshResponse{}

			if err = json.Unmarshal(got, &userRefreshResponse); err != nil {
				t.Errorf("%s - Failed - err=%s", helper.GetFunctionName(), err)
				return
			}

			if userRefreshResponse.EmailAddress != val.userRefreshRequest.EmailAddress {
				t.Errorf("%s - Failed - got=%s, want=%s", helper.GetFunctionName(), userRefreshResponse.EmailAddress, val.userRefreshRequest.EmailAddress)
				return
			}

			if userRefreshResponse.Success != val.want {
				t.Errorf("%s - Failed - got=%v, want=%v", helper.GetFunctionName(), userRefreshResponse.Success, val.want)
				return
			}
		})
	}
}
