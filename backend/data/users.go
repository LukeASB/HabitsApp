package data

import "time"

type UserAuth struct {
	EmailAddress string `json:"EmailAddress"`
	Password     string `json:"Password"`
}

type RegisterUser struct {
	Success bool     `json:"Succcess"`
	User    UserData `json:"User"`
}

type UserLoggedIn struct {
	Success     bool      `json:"Succcess"`
	User        UserData  `json:"User"`
	AccessToken string    `json:"AccessToken"`
	LoggedInAt  time.Time `json:"LoggedInAt"`
}

type UserLoggedOutResponse struct {
	Success      bool      `json:"Succcess"`
	UserID       string    `json:"UserID"`
	EmailAddress string    `json:"EmailAddress"`
	LoggedOutAt  time.Time `json:"LoggedOutAt"`
}

type UserLoggedOutRequest struct {
	UserID       string `json:"UserID"`
	EmailAddress string `json:"EmailAddress"`
}

type UserData struct {
	UserID       string    `json:"UserID"`
	Password     string    `json:"Password"`
	FirstName    string    `json:"FirstName"`
	LastName     string    `json:"LastName"`
	EmailAddress string    `json:"EmailAddress"`
	CreatedAt    time.Time `json:"CreatedAt"`
	LastLogin    time.Time `json:"LastLogin"`
	IsLoggedIn   bool      `json:"IsLoggedIn"`
}

type UserSession struct {
	ID           string    `json:"_id"`
	UserID       string    `json:"UserID"`
	RefreshToken string    `json:"RefreshToken"`
	Device       string    `json:"Device"`
	IPAddress    string    `json:"IpAddress"`
	CreatedAt    time.Time `json:"CreatedAt"`
}
