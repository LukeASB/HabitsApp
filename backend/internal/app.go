package internal

import (
	"dohabits/controller"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/middleware"
	"dohabits/middleware/session"
)

type App struct {
	authController   controller.IAuthController
	habitsController controller.IHabitsController
	database         db.IDB
	middleware       middleware.IMiddleware
	logger           logger.ILogger
	apiName          string
	apiVersion       string
	appVersion       string
	port             string
	jwtTokens        session.IJSONWebToken
}

type IApp interface {
	GetAuthController() controller.IAuthController
	GetHabitsController() controller.IHabitsController
	GetDB() db.IDB
	GetMiddleware() middleware.IMiddleware
	GetLogger() logger.ILogger
	GetAPIName() string
	GetAPIVersion() string
	GetAppVersion() string
	GetPort() string
}

func NewApp(
	authController *controller.AuthController,
	habitsController *controller.HabitsController,
	db db.IDB,
	middleware middleware.IMiddleware,
	logger *logger.Logger,
	apiName string,
	apiVersion string,
	appVersion string,
	port string,
	jwtTokens *session.JSONWebToken,
) *App {
	return &App{
		authController:   authController,
		habitsController: habitsController,
		database:         db,
		middleware:       middleware,
		logger:           logger,
		apiName:          apiName,
		apiVersion:       apiVersion,
		appVersion:       appVersion,
		port:             port,
		jwtTokens:        jwtTokens,
	}
}

func (a *App) GetAuthController() controller.IAuthController {
	return a.authController
}

func (a *App) GetHabitsController() controller.IHabitsController {
	return a.habitsController
}

func (a *App) GetDB() db.IDB {
	return a.database
}

func (a *App) GetMiddleware() middleware.IMiddleware {
	return a.middleware
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
