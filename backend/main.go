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
		os.Exit(1)
	}

	logger := &logger.Logger{}
	db := &db.MyMockDB{}
	m := &model.HabitsModel{}
	v := &view.HabitsView{}
	c := controller.NewHabitsController(logger)
	apiName := os.Getenv("API_NAME")
	apiVersion := os.Getenv("API_VERSION")
	appVersion := os.Getenv("APP_VERSION")
	port := os.Getenv("PORT")

	if err := db.Connect(logger); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	logger.SetVerbosity(2) // Will be set via env variable

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
	App.GetDB().Disconnect(App.GetLogger())
}
