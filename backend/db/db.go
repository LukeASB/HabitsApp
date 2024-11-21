package db

type IDB interface {
	Connect() error
	Disconnect() error
	RegisterUserHandler(value interface{}) (interface{}, error)
	LoginUser(value interface{}) error
	LogoutUser(value interface{}) error
	GetUserDetails(value interface{}) (interface{}, error)
	CreateHabitsHandler(value interface{}) error
	RetrieveAllHabitsHandler() (interface{}, error)
	RetrieveHabitsHandler(id string) (interface{}, error)
	UpdateHabitsHandler(id string, value interface{}) error
	DeleteHabitsHandler(id string) error
}
