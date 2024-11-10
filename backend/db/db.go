package db

type IDB interface {
	Connect() error
	Disconnect() error
	CreateHandler(value interface{}) error
	RetrieveAllHandler() (interface{}, error)
	RetrieveHandler(id string) (interface{}, error)
	UpdateHandler(id string, value interface{}) error
	DeleteHandler(id string) error
}
