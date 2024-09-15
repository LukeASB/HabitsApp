package internal

import (
	"dohabits/controller"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/model"
	"dohabits/view"
)

type App struct {
	model      model.IHabitsModel
	view       view.IHabitsView
	controller controller.IHabitsController
	database   db.IDB
	logger     logger.ILogger
	apiName    string
	apiVersion string
	appVersion string
	port       string
}

type IApp interface {
	GetHabitsModel() model.IHabitsModel
	GetView() view.IHabitsView
	GetController() controller.IHabitsController
	GetDB() db.IDB
	GetLogger() logger.ILogger
	GetAPIName() string
	GetAPIVersion() string
	GetAppVersion() string
	GetPort() string
}

// Use NewApp to declare so you can make the above struct private encapsulated
func NewApp(
	m *model.HabitsModel,
	v *view.HabitsView,
	c *controller.HabitsController,
	db db.IDB,
	logger *logger.Logger,
	apiName string,
	apiVersion string,
	appVersion string,
	port string) *App {
	return &App{
		model:      m,
		view:       v,
		controller: c,
		database:   db,
		logger:     logger,
		apiName:    apiName,
		apiVersion: apiVersion,
		appVersion: appVersion,
		port:       port,
	}
}

func (a *App) GetHabitsModel() model.IHabitsModel {
	return a.model
}

func (a *App) GetView() view.IHabitsView {
	return a.view
}

func (a *App) GetController() controller.IHabitsController {
	return a.controller
}

func (a *App) GetDB() db.IDB {
	return a.database
}

func (a *App) GetLogger() logger.ILogger {
	return a.logger
}

func (a *App) GetAPIName() string {
	return a.apiName
}

func (a *App) GetAPIVersion() string {
	return a.apiVersion
}

func (a *App) GetAppVersion() string {
	return a.appVersion
}

func (a *App) GetPort() string {
	return a.port
}
