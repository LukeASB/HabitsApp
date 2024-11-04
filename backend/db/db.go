package db

type IDB interface {
	Connect() error
	Disconnect() error
	Create(value interface{}) error
	RetrieveAll() (interface{}, error)
	Retrieve(id string) (interface{}, error)
	Update(id string, value interface{}) error
	Delete(id string) error
}
