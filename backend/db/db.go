package db

type IDB interface {
	Connect() error
	Disconnect() error
	RegisterUserHandler(value interface{}) (interface{}, error)
	LoginUser(value interface{}) error
	LogoutUser(value interface{}) error
	GetUserDetails(value interface{}) (interface{}, error)
	CreateHabitsHandler(userId string, value interface{}) error
	RetrieveAllHabitsHandler(userId string) (interface{}, error)
	RetrieveHabitsHandler(userId, habitId string) (interface{}, error)
	UpdateHabitsHandler(userId, habitId string, value interface{}) error
	DeleteHabitsHandler(userId, habitId string) error
}
