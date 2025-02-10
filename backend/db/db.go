package db

import "dohabits/data"

type IDB interface {
	Connect() error
	Disconnect() error
	RegisterUserHandler(value interface{}) (interface{}, error)
	LoginUser(value interface{}) error
	LogoutUser(value interface{}) error
	RetrieveUserSession(value interface{}, userID string) (string, error)
	RetrieveUserDetails(value interface{}) (interface{}, error)
	CreateHabitsHandler(userId string, value interface{}) (*data.NewHabitResponse, error)
	RetrieveAllHabitsHandler(userId string) (interface{}, error)
	RetrieveHabitsHandler(userId, habitId string) (interface{}, error)
	UpdateHabitsHandler(userId, habitId string, value interface{}) error
	UpdateAllHabitsHandler(userId string, value interface{}) error
	DeleteHabitsHandler(userId, habitId string) error
}
