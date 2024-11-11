package data

import "time"

type UserAuth struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserLoggedIn struct {
	Success    bool      `json:"Succcess"`
	Username   string    `json:"Username"`
	LoggedInAt time.Time `json:"LoggedInAt"`
}

type UserLoggedOut struct {
	Success     bool      `json:"Succcess"`
	Username    string    `json:"Username"`
	LoggedOutAt time.Time `json:"LoggedOutAt"`
}
