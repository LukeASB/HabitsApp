package controller

import (
	"dohabits/data"
	"dohabits/logger"
	"dohabits/middleware/session"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"net/http"
)

type HabitsController struct {
	opsChan     chan func()
	habitsModel model.IHabitsModel
	habitsView  view.IHabitsView
	logger      logger.ILogger
}

type IHabitsController interface {
	CreateHabitsHandler(w http.ResponseWriter, r *http.Request)
	RetrieveHabitsHandler(w http.ResponseWriter, r *http.Request)
	RetrieveAllHabitsHandler(w http.ResponseWriter, r *http.Request)
	UpdateHabitsHandler(w http.ResponseWriter, r *http.Request)
	DeleteHabitsHandler(w http.ResponseWriter, r *http.Request)
}

// Initialise the processOperations Goroutine
func NewHabitsController(habitsModel model.IHabitsModel, habitsView view.IHabitsView, logger logger.ILogger) *HabitsController {
	logger.InfoLog("habitsController.NewHabitsController")

	habitsController := &HabitsController{
		opsChan:     make(chan func(), 1),
		habitsModel: habitsModel,
		habitsView:  habitsView,
		logger:      logger,
	}

	go habitsController.manageOps(logger) // Run the Goroutine to handle operations.

	return habitsController
}

func (c *HabitsController) manageOps(logger logger.ILogger) {
	logger.InfoLog("habitsController.manageOps")
	// Wait and execute any function that get sent to opsChan
	for op := range c.opsChan {
		logger.DebugLog("habitsController.manageOps - exec func passed to channel")
		op() // Execute the function passed to the channel
	}
}

func (c *HabitsController) CreateHabitsHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		c.logger.InfoLog("habitsController.Create")

		claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

		if !ok {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims not found"),
			}
			return
		}

		username := claims.Username

		if username == "" {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims username is empty"),
			}
			return
		}

		if r.Body == nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - Body is empty"),
			}
			return
		}

		newHabit := data.NewHabit{}

		if err := json.NewDecoder(r.Body).Decode(&newHabit); err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - Erroring decoding newHabit JSON - err=%s", err),
			}
			return
		}

		if err := c.habitsModel.CreateHabitsHandler(username, newHabit); err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - model.Create - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.CreateHabitsHandler(newHabit)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - view.Create - err=%s", err),
			}
			return
		}

		resultChan <- struct {
			data []byte
			err  error
		}{
			data: result,
			err:  nil,
		}
	}

	res := <-resultChan

	if res.err != nil {
		c.logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveHabitsHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

		if !ok {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims not found"),
			}
			return
		}

		username := claims.Username

		if username == "" {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims username is empty"),
			}
			return
		}

		habitId := r.URL.Query().Get("habitId")

		if len(habitId) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - habitId query param is empty"),
			}
			return
		}

		c.logger.InfoLog(fmt.Sprintf("habitsController.Retrieve - email=%s, habitId=%s", username, habitId))

		habit, err := c.habitsModel.RetrieveHabitsHandler(username, habitId)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - model.Retrieve - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.RetrieveHabitsHandler(habit)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - view.Retrieve - err=%s", err),
			}
			return
		}

		resultChan <- struct {
			data []byte
			err  error
		}{
			data: result,
			err:  nil,
		}
	}

	res := <-resultChan

	if res.err != nil {
		c.logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Retrieve - Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveAllHabitsHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

		if !ok {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims not found"),
			}
			return
		}

		username := claims.Username

		if username == "" {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims username is empty"),
			}
			return
		}

		c.logger.InfoLog(fmt.Sprintf("habitsController.RetrieveAll = email=%s", username))

		habits, err := c.habitsModel.RetrieveAllHabitsHandler(username)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.RetrieveAll - model.RetrieveAll - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.RetrieveAllHabitsHandler(habits)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.RetrieveAll - view.RetrieveAll - err=%s", err),
			}
			return
		}

		resultChan <- struct {
			data []byte
			err  error
		}{
			data: result,
			err:  nil,
		}
	}

	res := <-resultChan

	if res.err != nil {
		c.logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.RetrieveAll - Error writing response: %s", err))
	}
}

func (c *HabitsController) UpdateHabitsHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

		if !ok {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims not found"),
			}
			return
		}

		username := claims.Username

		if username == "" {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims username is empty"),
			}
			return
		}

		habitId := r.URL.Query().Get("habitId")

		c.logger.InfoLog(fmt.Sprintf("habitsController.Update - email=%s, habitId=%s", username, habitId))

		if len(habitId) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Update - habitId query param is missing"),
			}
			return
		}

		updatedHabit := data.UpdateHabit{}

		err := json.NewDecoder(r.Body).Decode(&updatedHabit)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Update - Erroring decoding newHabit JSON - err=%s", err),
			}
			return
		}

		habit, err := c.habitsModel.RetrieveHabitsHandler(username, habitId)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Update - err=%s", err),
			}
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
			habit.CompletionDates = append(habit.CompletionDates, *updatedHabit.CompletionDates...)
		}

		err = c.habitsModel.UpdateHabitsHandler(username, habit, habitId)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - model.Update - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.UpdateHabitsHandler(habit)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Update - view.Update - err=%s", err),
			}
			return
		}

		resultChan <- struct {
			data []byte
			err  error
		}{
			data: result,
			err:  nil,
		}
	}

	res := <-resultChan

	if res.err != nil {
		c.logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Update - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Update - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Update - Error writing response: %s", err))
	}
}

func (c *HabitsController) DeleteHabitsHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		claims, ok := r.Context().Value(session.ClaimsKey).(*session.Claims)

		if !ok {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims not found"),
			}
			return
		}

		username := claims.Username

		if username == "" {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - JWT Token claims username is empty"),
			}
			return
		}

		habitId := r.URL.Query().Get("habitId")

		c.logger.InfoLog(fmt.Sprintf("habitsController.Delete - email=%s, habitId=%s", username, habitId))

		if len(habitId) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Delete - habitId query param is missing"),
			}
			return
		}

		err := c.habitsModel.DeleteHabitsHandler(username, habitId)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Delete - model.Delete - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.DeleteHabitsHandler()

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Delete - view.Delete - err=%s", err),
			}
			return
		}

		resultChan <- struct {
			data []byte
			err  error
		}{
			data: result,
			err:  nil,
		}
	}

	res := <-resultChan

	if res.err != nil {
		c.logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}
