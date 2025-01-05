package main

import (
	"dohabits/controller"
	"dohabits/db"
	"dohabits/internal"
	"dohabits/logger"
	"dohabits/middleware"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/routes"
	"dohabits/view"
	"fmt"
	"log"
	"net/http"
	"os"
)

var App *internal.App

func init() {
	if err := internal.LoadEnvVariables(); err != nil {
		log.Fatal(err)
	}

	logger := &logger.Logger{}

	if err := logger.SetVerbosity(os.Getenv("LOG_VERBOSITY")); err != nil {
		log.Fatal(err)
	}

	db := db.NewMongoDB(logger)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	jwtTokens := session.NewJWTTokens(os.Getenv("JWT_SECRET"))
	csrfTokens := session.NewCSRFToken(logger)

	authModel := model.NewAuthModel(logger, db)
	habitsModel := model.NewHabitsModel(logger, db)
	authView := view.NewAuthView(logger)
	habitsView := view.NewHabitsView(logger)
	authController := controller.NewAuthController(authModel, authView, jwtTokens, csrfTokens, logger)
	habitsController := controller.NewHabitsController(habitsModel, habitsView, logger)

	mw := middleware.NewMiddleware(jwtTokens, csrfTokens, logger)
	apiName := os.Getenv("API_NAME")
	apiVersion := os.Getenv("API_VERSION")
	appVersion := os.Getenv("APP_VERSION")
	port := os.Getenv("PORT")

	App = internal.NewApp(authController, habitsController, db, mw, logger, apiName, apiVersion, appVersion, port, jwtTokens)

	App.GetLogger().DebugLog(fmt.Sprintf("main.init() - %s loaded successfully. App Version = %s, API Version = %s", App.GetAPIName(), App.GetAppVersion(), App.GetAPIVersion()))
}

func main() {
	App.GetLogger().DebugLog("main.main - Executed")
	routes.SetUpRoutes(App)

	defer cleanup()

	App.GetLogger().InfoLog(fmt.Sprintf("Listening on port: :%s\n", App.GetPort()))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", App.GetPort()), nil); err != nil {
		App.GetLogger().ErrorLog("The Habits App has Exploded: 💣")
		os.Exit(1)
	}
}

func cleanup() {
	App.GetLogger().DebugLog("main.cleanup - Executed")
	App.GetDB().Disconnect()
}
