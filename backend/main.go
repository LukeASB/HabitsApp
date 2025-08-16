package main

import (
	"dohabits/controller"
	"dohabits/db"
	"dohabits/helper"
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
	"strconv"
)

var App *internal.App

func init() {
	if err := internal.LoadEnvVariables(); err != nil {
		log.Fatal(err)
	}

	verbosity, err := strconv.Atoi(os.Getenv("LOG_VERBOSITY"))

	if err != nil {
		log.Fatalf("Log Verbosity must be an integer: err=%s", err)
	}

	logger := logger.NewLogger(verbosity)

	var database interface{}

	if os.Getenv("ENVIRONMENT") == "DEV" {
		logger.InfoLog(helper.GetFunctionName(), "Connected to Mock DB dataset.")
		database = db.NewMockDB(logger)
	} else {
		logger.InfoLog(helper.GetFunctionName(), "Connected to Production DB.")
		database = db.NewMongoDB(logger)
	}

	db, ok := database.(db.IDB)

	if !ok {
		logger.ErrorLog(helper.GetFunctionName(), "db is not of type IDB.")
		log.Fatal("db is not of type IDB.")
	}

	if err := db.Connect(); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("An error occured when connecting to the DB. err=%s", err))
		log.Fatal(err)
	}

	jwtTokens := session.NewJSONWebToken(os.Getenv("JWT_SECRET"), db)
	csrfTokens := session.NewCSRFToken(os.Getenv("JWT_SECRET"), logger)

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

	App.GetLogger().DebugLog(helper.GetFunctionName(), fmt.Sprintf("%s loaded successfully. App Version = %s, API Version = %s", App.GetAPIName(), App.GetAppVersion(), App.GetAPIVersion()))
}

func main() {
	App.GetLogger().DebugLog(helper.GetFunctionName(), "Executed")
	routes.SetUpRoutes(App)

	defer cleanup()

	App.GetLogger().InfoLog(helper.GetFunctionName(), fmt.Sprintf("Listening on port: :%s\n", App.GetPort()))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", App.GetPort()), nil); err != nil {
		App.GetLogger().ErrorLog(helper.GetFunctionName(), "The Habits App has Exploded: 💣")
		os.Exit(1)
	}
}

func cleanup() {
	App.GetLogger().DebugLog(helper.GetFunctionName(), "Executed")
	App.GetDB().Disconnect()
}
