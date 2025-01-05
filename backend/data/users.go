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
	UserID       string    `json:"UserID" bson:"_id"`
	Password     string    `json:"Password" bson:"Password"`
	FirstName    string    `json:"FirstName" bson:"FirstName"`
	LastName     string    `json:"LastName" bson:"LastName"`
	EmailAddress string    `json:"EmailAddress" bson:"EmailAddress"`
	CreatedAt    time.Time `json:"CreatedAt" bson:"CreatedAt"`
	LastLogin    time.Time `json:"LastLogin" bson:"LastLogin"`
	IsLoggedIn   bool      `json:"IsLoggedIn" bson:"IsLoggedIn"`
}

type UserSession struct {
	UserID       string    `json:"UserID" bson:"_id"`
	RefreshToken string    `json:"RefreshToken" bson:"RefreshToken"`
	Device       string    `json:"Device" bson:"Device"`
	IPAddress    string    `json:"IpAddress" bson:"IpAddress"`
	CreatedAt    time.Time `json:"CreatedAt" bson:"CreatedAt"`
}
