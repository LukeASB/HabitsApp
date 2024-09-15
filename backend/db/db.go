package db

import "dohabits/logger"

type IDB interface {
	Connect(logger logger.ILogger) error
	Disconnect(logger logger.ILogger) error
	Create(logger logger.ILogger, value interface{}) error
	RetrieveAll(logger logger.ILogger) (interface{}, error)
	Retrieve(logger logger.ILogger, id string) (interface{}, error)
	Update(logger logger.ILogger, id string, value interface{}) error
	Delete(logger logger.ILogger, id string) error
}
