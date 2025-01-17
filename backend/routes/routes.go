package routes

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/internal"
	"fmt"
	"net/http"
)

func SetUpRoutes(app internal.IApp) {
	endpoint := fmt.Sprintf("%s/%s", app.GetAPIName(), app.GetAPIVersion())

	app.GetLogger().InfoLog(helper.GetFunctionName(), fmt.Sprintf("endpoint = %s", endpoint))

	http.HandleFunc(fmt.Sprintf("/%s/register", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().RegisterUserHandler, data.Middleware{HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/login", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().LoginHandler, data.Middleware{HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/logout", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().LogoutHandler, data.Middleware{IsProtected: true, CSRFRequired: true, HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/refresh", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().RefreshHandler, data.Middleware{HTTPMethod: http.MethodPost}))

	http.HandleFunc(fmt.Sprintf("/%s/createhabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().CreateHabitsHandler, data.Middleware{IsProtected: true, CSRFRequired: true, HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/retrievehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().RetrieveHabitsHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodGet}))
	http.HandleFunc(fmt.Sprintf("/%s/retrievehabits", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().RetrieveAllHabitsHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodGet}))
	http.HandleFunc(fmt.Sprintf("/%s/updatehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().UpdateHabitsHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodPut}))
	http.HandleFunc(fmt.Sprintf("/%s/deletehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().DeleteHabitsHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodDelete}))
}
