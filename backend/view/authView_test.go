package view

import (
	"bytes"
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"testing"
	"time"
)

func TestRegisterUserHandler(t *testing.T) {
	logger := &logger.Logger{}
	authView := NewAuthView(logger)

	fixedTime := time.Date(2023, time.November, 19, 12, 0, 0, 0, time.UTC)

	testCases := []struct {
		name             string
		registerUserData *data.RegisterUserData
		want             *data.RegisterUserResponse
	}{
		{
			name: "Test Register Response",
			registerUserData: &data.RegisterUserData{
				Success: true,
				User: data.UserData{
					UserID:       "1",
					Password:     "password",
					FirstName:    "first",
					LastName:     "last",
					EmailAddress: "email@email.com",
					CreatedAt:    fixedTime,
					LastLogin:    fixedTime,
					IsLoggedIn:   true,
				},
			},
			want: &data.RegisterUserResponse{
				Success: true,
				User: data.UserDataResponse{
					FirstName:    "first",
					LastName:     "last",
					EmailAddress: "email@email.com",
					CreatedAt:    fixedTime,
				},
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledRegisterUserResponse, err := json.Marshal(val.want)

			if err != nil {
				t.Errorf("TestRegisterUserHandler - Fail err: %s", err)
			}

			got, err := authView.RegisterUserHandler(val.registerUserData)

			if err != nil {
				t.Errorf("TestRegisterUserHandler - Fail err: %s", err)
			}

			if !bytes.Equal(marshalledRegisterUserResponse, got) {
				t.Errorf("TestRegisterUserHandler - Fail want doesn't match got")
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	logger := &logger.Logger{}
	authView := NewAuthView(logger)

	fixedTime := time.Date(2023, time.November, 19, 12, 0, 0, 0, time.UTC)

	testCases := []struct {
		name             string
		userLoggedInData *data.UserLoggedInData
		want             *data.UserLoggedInResponse
	}{
		{
			name: "Test Login Response",
			userLoggedInData: &data.UserLoggedInData{
				Success: true,
				User: data.UserData{
					UserID:       "1",
					Password:     "password",
					FirstName:    "first",
					LastName:     "last",
					EmailAddress: "email@email.com",
					CreatedAt:    fixedTime,
					LastLogin:    fixedTime,
					IsLoggedIn:   true,
				},
				AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5kb2UxQGV4YW1wbGUuY29tIiwiZXhwIjoxNzMyMjU5NjUzfQ.vu2Vv_2z--i3p8TLYIHRmyKX9xjyICr_esCGrGYs2Es",
				LoggedInAt:  fixedTime,
			},
			want: &data.UserLoggedInResponse{
				Success: true,
				User: data.UserDataResponse{
					FirstName:    "first",
					LastName:     "last",
					EmailAddress: "email@email.com",
					CreatedAt:    fixedTime,
				},
				LoggedInAt: fixedTime,
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledUserLoggedInResponse, err := json.Marshal(val.want)

			if err != nil {
				t.Errorf("TestLoginHandler - Fail err: %s", err)
			}

			got, err := authView.LoginHandler(val.userLoggedInData)

			if err != nil {
				t.Errorf("TestLoginHandler - Fail err: %s", err)
			}

			if !bytes.Equal(marshalledUserLoggedInResponse, got) {
				t.Errorf("TestLoginHandler - Fail want doesn't match got")
			}
		})
	}
}

func TestRefreshHandler(t *testing.T) {
	logger := &logger.Logger{}
	authView := NewAuthView(logger)

	testCases := []struct {
		name               string
		userRefreshRequest *data.UserRefreshRequest
		want               *data.UserRefreshResponse
	}{
		{
			name: "Test User Refresh Response",
			userRefreshRequest: &data.UserRefreshRequest{
				EmailAddress: "test@email.com",
			},
			want: &data.UserRefreshResponse{
				Success:      true,
				EmailAddress: "test@email.com",
			},
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			marshalledUserRefreshResponse, err := json.Marshal(val.want)

			if err != nil {
				t.Errorf("TestRefreshHandler - Fail err: %s", err)
			}

			got, err := authView.RefreshHandler(val.userRefreshRequest)

			if err != nil {
				t.Errorf("TestRefreshHandler - Fail err: %s", err)
			}

			if !bytes.Equal(marshalledUserRefreshResponse, got) {
				t.Errorf("TestRefreshHandler - Fail want doesn't match got")
			}
		})
	}
}
