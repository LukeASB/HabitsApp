package controller

import (
	"bytes"
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
	db := db.NewDB(logger)
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
				Password:     "secretPassword012!",
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
				t.Errorf("TestLoginHandler - HTTP Status Code = %d", status)
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
	db := db.NewDB(logger)
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
				Password:     "secretPassword012!",
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
	db := db.NewDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name                 string
		want                 bool
		userLoggedOutRequest *data.UserLoggedOutRequest
	}{
		{
			name: "Successfully logout",
			want: true,
			userLoggedOutRequest: &data.UserLoggedOutRequest{
				UserID:       "4",
				EmailAddress: "john.loggedin@example.com",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledLoggedOutRequest, err := json.Marshal(val.userLoggedOutRequest)

			if err != nil {
				t.Errorf("TestLogoutHandler err: %s", err)
				return
			}

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/%s/logout", endpoint), io.NopCloser(bytes.NewBuffer(marshalledLoggedOutRequest)))
			w := httptest.NewRecorder()

			authController.LogoutHandler(w, req)

			if status := w.Code; status == http.StatusInternalServerError {
				t.Errorf("TestLogoutHandler - HTTP Status Code = %d", status)
				return
			}

			res := w.Result()

			defer res.Body.Close()

			got, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("TestLogoutHandler err: %s", err)
				return
			}

			userLoggedOutResponse := &data.UserLoggedOutResponse{}

			err = json.Unmarshal(got, userLoggedOutResponse)

			if err != nil {
				t.Errorf("TestLogoutHandler err: %s", err)
				return
			}

			if userLoggedOutResponse.UserID != val.userLoggedOutRequest.UserID {
				t.Errorf("TestLoginHandler - userId does not match logged out user. got=%s want=%s", userLoggedOutResponse.UserID, val.userLoggedOutRequest.UserID)
				return
			}

			if userLoggedOutResponse.EmailAddress != val.userLoggedOutRequest.EmailAddress {
				t.Errorf("TestLoginHandler - email does not match logged out  user. got=%s want=%s", userLoggedOutResponse.EmailAddress, val.userLoggedOutRequest.EmailAddress)
				return
			}

			if userLoggedOutResponse.Success != val.want {
				t.Errorf("TestRegisterUser Failed - got=%v, want=%v", userLoggedOutResponse.Success, val.want)
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewDB(logger)
	jwtTokensMock := session.NewMockJWTTokens("secretJwt")
	csrfTokenMock := session.NewMockCSRFToken(logger)
	authModel := model.NewAuthModel(logger, db)
	authView := view.NewAuthView(logger)
	authController := NewAuthController(authModel, authView, jwtTokensMock, csrfTokenMock, logger)

	endpoint := fmt.Sprintf("%s/%s", os.Getenv("API_NAME"), os.Getenv("API_VERSION"))

	testCases := []struct {
		name               string
		want               string
		userRefreshRequest *data.UserRefreshRequest
	}{
		{
			name: "Successfully Get a Refresh Token",
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es",
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

			if userRefreshResponse.AccessToken != val.want {
				t.Errorf("TestRefreshHandler Failed - got=%v, want=%v", userRefreshResponse.AccessToken, val.want)
			}
		})
	}
}
