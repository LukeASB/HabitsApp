package controller

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type HabitsController struct {
	opsChan     chan func()
	habitsModel model.IHabitsModel
	habitsView  view.IHabitsView
	logger      logger.ILogger
	mx          sync.RWMutex
}

type IHabitsController interface {
	CreateHabitsHandler(w http.ResponseWriter, r *http.Request)
	RetrieveHabitsHandler(w http.ResponseWriter, r *http.Request)
	RetrieveAllHabitsHandler(w http.ResponseWriter, r *http.Request)
	UpdateHabitsHandler(w http.ResponseWriter, r *http.Request)
	UpdateAllHabitsHandler(w http.ResponseWriter, r *http.Request)
	DeleteHabitsHandler(w http.ResponseWriter, r *http.Request)
}

// Initialise the processOperations Goroutine
func NewHabitsController(habitsModel model.IHabitsModel, habitsView view.IHabitsView, logger logger.ILogger) *HabitsController {
	logger.InfoLog(helper.GetFunctionName(), "")

	habitsController := &HabitsController{
		opsChan:     make(chan func(), 1),
		habitsModel: habitsModel,
		habitsView:  habitsView,
		logger:      logger,
		mx:          sync.RWMutex{},
	}

	return habitsController
}

func (c *HabitsController) CreateHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.logger.InfoLog(helper.GetFunctionName(), "")

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.Body == nil {
		c.logger.ErrorLog(helper.GetFunctionName(), "Body is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newHabit := data.NewHabit{}

	if err := json.NewDecoder(r.Body).Decode(&newHabit); err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error decoding JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newHabitResponse, err := c.habitsModel.CreateHabitsHandler(username, newHabit)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.CreateHabitsHandler(newHabitResponse)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	habitId := r.URL.Query().Get("habitId")

	if len(habitId) == 0 {
		c.logger.ErrorLog(helper.GetFunctionName(), "habitId query param is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("email=%s, habitId=%s", username, habitId))

	habit, err := c.habitsModel.RetrieveHabitsHandler(username, habitId)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.RetrieveHabitsHandler(habit)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveAllHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("email=%s", username))

	habits, err := c.habitsModel.RetrieveAllHabitsHandler(username)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.RetrieveAllHabitsHandler(habits)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}

func (c *HabitsController) UpdateHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.Lock()
	defer c.mx.Unlock()

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	updatedHabit := data.UpdateHabit{}

	err := json.NewDecoder(r.Body).Decode(&updatedHabit)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error decoding newHabit JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if updatedHabit.HabitID == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "habitId is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("email=%s, habitId=%s", username, updatedHabit.HabitID))

	habit, err := c.habitsModel.RetrieveHabitsHandler(username, updatedHabit.HabitID)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if updatedHabit.Name != nil {
		habit.Name = *updatedHabit.Name
	}

	if updatedHabit.Days != nil {
		habit.Days = *updatedHabit.Days
	}

	if updatedHabit.DaysTarget != nil {
		habit.DaysTarget = *updatedHabit.DaysTarget
	}

	if updatedHabit.CompletionDates != nil {
		habit.CompletionDates = *updatedHabit.CompletionDates
	}

	err = c.habitsModel.UpdateHabitsHandler(username, habit, updatedHabit.HabitID)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.UpdateHabitsHandler(habit)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}

func (c *HabitsController) UpdateAllHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.Lock()
	defer c.mx.Unlock()

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	updatedHabit := []data.UpdateHabit{}

	err := json.NewDecoder(r.Body).Decode(&updatedHabit)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error decoding JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userHabits, err := c.habitsModel.RetrieveAllHabitsHandler(username)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for i, habit := range userHabits {
		for _, updatedHabit := range updatedHabit {
			if habit.HabitID == updatedHabit.HabitID {
				if updatedHabit.Name != nil {
					userHabits[i].Name = *updatedHabit.Name
				}

				if updatedHabit.Days != nil {
					userHabits[i].Days = *updatedHabit.Days
				}

				if updatedHabit.DaysTarget != nil {
					userHabits[i].DaysTarget = *updatedHabit.DaysTarget
				}

				if updatedHabit.CompletionDates != nil {
					userHabits[i].CompletionDates = *updatedHabit.CompletionDates
				}
			}
		}
	}

	err = c.habitsModel.UpdateAllHabitsHandler(username, &userHabits)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.UpdateAllHabitsHandler(&userHabits)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}

func (c *HabitsController) DeleteHabitsHandler(w http.ResponseWriter, r *http.Request) {
	c.mx.Lock()
	defer c.mx.Unlock()

	claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

	if !ok {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims not found")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	username := claims.Username

	if username == "" {
		c.logger.ErrorLog(helper.GetFunctionName(), "JWT Token claims username is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	habitId := r.URL.Query().Get("habitId")

	c.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("email=%s, habitId=%s", username, habitId))

	if len(habitId) == 0 {
		c.logger.ErrorLog(helper.GetFunctionName(), "habitId is empty")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err := c.habitsModel.DeleteHabitsHandler(username, habitId)

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := c.habitsView.DeleteHabitsHandler()

	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("Writing response: %s", string(result)))
	numOfBytes, err := w.Write(result)
	c.logger.DebugLog(helper.GetFunctionName(), fmt.Sprintf("w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error writing response: %s", err))
	}
}
