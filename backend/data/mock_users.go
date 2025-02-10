package data

import "time"

// Hash the passwords - Hashed with "golang.org/x/crypto/bcrypt"
var MockUsers = []UserData{
	{
		UserID:       "1",
		Password:     "$2a$10$hLa8z.sjayeZNWNAcise5.VKvgkcftE8z0n/mze7O8zXJwOQ9M5tW", // 1secret?Password
		FirstName:    "John",
		LastName:     "Doe",
		EmailAddress: "johndoe1@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
	},
	{
		UserID:       "2",
		Password:     "$2a$10$hLa8z.sjayeZNWNAcise5.VKvgkcftE8z0n/mze7O8zXJwOQ9M5tW", // 1secret?Password
		FirstName:    "Jane",
		LastName:     "Smith",
		EmailAddress: "janesmith@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
	},
	{
		UserID:       "3",
		Password:     "$2a$10$hLa8z.sjayeZNWNAcise5.VKvgkcftE8z0n/mze7O8zXJwOQ9M5tW", // 1secret?Password
		FirstName:    "Alice",
		LastName:     "Johnson",
		EmailAddress: "alicejohnson@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
	},
	{
		UserID:       "4",
		Password:     "$2a$10$hLa8z.sjayeZNWNAcise5.VKvgkcftE8z0n/mze7O8zXJwOQ9M5tW", // 1secret?Password
		FirstName:    "John",
		LastName:     "LoggedIn",
		EmailAddress: "john.loggedin@example.com",
		CreatedAt:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
		LastLogin:    time.Date(2024, time.October, 10, 9, 0, 0, 0, time.UTC),
	},
}

var MockUserSession = []UserSession{}
