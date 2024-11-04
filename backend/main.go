package main

import (
	"dohabits/controller"
	"dohabits/db"
	"dohabits/internal"
	"dohabits/logger"
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

	db := db.NewDB(logger)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	m := model.NewHabitsModel(logger, db)
	v := view.NewHabitsView(logger)
	c := controller.NewHabitsController(logger)
	apiName := os.Getenv("API_NAME")
	apiVersion := os.Getenv("API_VERSION")
	appVersion := os.Getenv("APP_VERSION")
	port := os.Getenv("PORT")

	App = internal.NewApp(m, v, c, db, logger, apiName, apiVersion, appVersion, port)

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
