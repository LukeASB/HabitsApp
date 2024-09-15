package routes

import (
	"dohabits/internal"
	"fmt"
	"net/http"
)

func SetUpRoutes(app internal.IApp) {
	endpoint := fmt.Sprintf("%s/%s", app.GetAPIName(), app.GetAPIVersion())

	app.GetLogger().InfoLog(fmt.Sprintf("routes.SetUpRoutes() - endpoint = %s", endpoint))

	http.HandleFunc(fmt.Sprintf("/%s/CreateHabit", endpoint), func(w http.ResponseWriter, r *http.Request) {
		app.GetController().Create(w, r, app.GetHabitsModel(), app.GetView(), app.GetDB(), app.GetLogger())
	})

	http.HandleFunc(fmt.Sprintf("/%s/RetrieveHabit", endpoint), func(w http.ResponseWriter, r *http.Request) {
		app.GetController().Retrieve(w, r, app.GetHabitsModel(), app.GetView(), app.GetDB(), app.GetLogger())
	})

	http.HandleFunc(fmt.Sprintf("/%s/RetrieveAllHabits", endpoint), func(w http.ResponseWriter, r *http.Request) {
		app.GetController().RetrieveAll(w, r, app.GetHabitsModel(), app.GetView(), app.GetDB(), app.GetLogger())
	})

	http.HandleFunc(fmt.Sprintf("/%s/UpdateHabit", endpoint), func(w http.ResponseWriter, r *http.Request) {
		app.GetController().Update(w, r, app.GetHabitsModel(), app.GetView(), app.GetDB(), app.GetLogger())
	})

	http.HandleFunc(fmt.Sprintf("/%s/DeleteHabit", endpoint), func(w http.ResponseWriter, r *http.Request) {
		app.GetController().Delete(w, r, app.GetHabitsModel(), app.GetView(), app.GetDB(), app.GetLogger())
	})
}
