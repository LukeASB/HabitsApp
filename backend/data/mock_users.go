package data

import "time"

// Hash the passwords - Hashed with "golang.org/x/crypto/bcrypt"
var MockUsers = []UserData{
	{
		UserID:       "1",
		Password:     "$2a$10$jJ75MWcPyEOnwBUsg4g9I.wLdpPagwsNtb0i.mhwpNYqdt8mpmdI.", // secretPassword012!
		FirstName:    "John",
		LastName:     "Doe",
		EmailAddress: "johndoe1@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		IsLoggedIn:   false,
	},
	{
		UserID:       "2",
		Password:     "$2a$10$jJ75MWcPyEOnwBUsg4g9I.wLdpPagwsNtb0i.mhwpNYqdt8mpmdI.", // secretPassword012!
		FirstName:    "Jane",
		LastName:     "Smith",
		EmailAddress: "janesmith@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		IsLoggedIn:   false,
	},
	{
		UserID:       "3",
		Password:     "$2a$10$jJ75MWcPyEOnwBUsg4g9I.wLdpPagwsNtb0i.mhwpNYqdt8mpmdI.", // secretPassword012!
		FirstName:    "Alice",
		LastName:     "Johnson",
		EmailAddress: "alicejohnson@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		IsLoggedIn:   false,
	},
}

var MockUserSession = []UserSession{}
