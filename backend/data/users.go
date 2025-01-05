package data

import "time"

type UserAuth struct {
	UserID       string `json:"UserID"`
	EmailAddress string `json:"EmailAddress"`
	Password     string `json:"Password"`
}

type UserRefreshRequest struct {
	EmailAddress string `json:"EmailAddress"`
}

type UserRefreshResponse struct {
	Success      bool   `json:"Success"`
	EmailAddress string `json:"EmailAddress"`
}
type RegisterUserRequest struct {
	EmailAddress string `json:"EmailAddress"`
	Password     string `json:"Password"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
}

type RegisterUserData struct {
	Success bool     `json:"Success"`
	User    UserData `json:"User"`
}

type RegisterUserResponse struct {
	Success bool             `json:"Success"`
	User    UserDataResponse `json:"User"`
}

type UserDataResponse struct {
	FirstName    string    `json:"FirstName"`
	LastName     string    `json:"LastName"`
	EmailAddress string    `json:"EmailAddress"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

type UserLoggedInData struct {
	Success     bool      `json:"Success"`
	User        UserData  `json:"User"`
	AccessToken string    `json:"AccessToken"`
	LoggedInAt  time.Time `json:"LoggedInAt"`
}

type UserLoggedInResponse struct {
	Success    bool             `json:"Success"`
	User       UserDataResponse `json:"User"`
	LoggedInAt time.Time        `json:"LoggedInAt"`
}

type UserLoggedOutRequest struct {
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
