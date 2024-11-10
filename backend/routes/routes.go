package routes

import (
	"dohabits/data"
	"dohabits/internal"
	"fmt"
	"net/http"
)

func SetUpRoutes(app internal.IApp) {
	endpoint := fmt.Sprintf("%s/%s", app.GetAPIName(), app.GetAPIVersion())

	app.GetLogger().InfoLog(fmt.Sprintf("routes.SetUpRoutes() - endpoint = %s", endpoint))

	http.HandleFunc(fmt.Sprintf("/%s/login", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().LoginHandler, data.Middleware{HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/logout", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().LogoutHandler, data.Middleware{HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/refresh", endpoint), app.GetMiddleware().MiddlewareList(app.GetAuthController().RefreshHandler, data.Middleware{HTTPMethod: http.MethodPost}))

	http.HandleFunc(fmt.Sprintf("/%s/createhabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().CreateHandler, data.Middleware{IsProtected: true, CSRFRequired: true, HTTPMethod: http.MethodPost}))
	http.HandleFunc(fmt.Sprintf("/%s/retrievehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().RetrieveHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodGet}))
	http.HandleFunc(fmt.Sprintf("/%s/retrievehabits", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().RetrieveAllHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodGet}))
	http.HandleFunc(fmt.Sprintf("/%s/updatehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().UpdateHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodPut}))
	http.HandleFunc(fmt.Sprintf("/%s/deletehabit", endpoint), app.GetMiddleware().MiddlewareList(app.GetHabitsController().DeleteHandler, data.Middleware{IsProtected: true, HTTPMethod: http.MethodDelete}))
}
