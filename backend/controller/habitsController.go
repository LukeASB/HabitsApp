package controller

import (
	"dohabits/data"
	"dohabits/logger"
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
	CreateHandler(w http.ResponseWriter, r *http.Request)
	RetrieveHandler(w http.ResponseWriter, r *http.Request)
	RetrieveAllHandler(w http.ResponseWriter, r *http.Request)
	UpdateHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
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

func (c *HabitsController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		c.logger.InfoLog("habitsController.Create")

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

		err := json.NewDecoder(r.Body).Decode(&newHabit)

		if err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - Erroring decoding newHabit JSON - err=%s", err),
			}
			return
		}

		if err := c.habitsModel.CreateHandler(newHabit); err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - model.Create - err=%s", err),
			}
			return
		}

		result, err := c.habitsView.CreateHandler(newHabit)

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
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		id := r.URL.Query().Get("id")

		c.logger.InfoLog(fmt.Sprintf("habitsController.Retrieve - id=%s", id))

		if len(id) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Retrieve - id query param is empty"),
			}
			return
		}

		habit, err := c.habitsModel.RetrieveHandler(id)

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

		result, err := c.habitsView.RetrieveHandler(habit)

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
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Retrieve - Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveAllHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		c.logger.InfoLog("habitsController.RetrieveAll")

		habits, err := c.habitsModel.RetrieveAllHandler()

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

		result, err := c.habitsView.RetrieveAllHandler(habits)

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
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.RetrieveAll - Error writing response: %s", err))
	}
}

func (c *HabitsController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		id := r.URL.Query().Get("id")

		c.logger.InfoLog(fmt.Sprintf("habitsController.Update - id=%s", id))

		if len(id) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Update - id query param is missing"),
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

		habit, err := c.habitsModel.RetrieveHandler(id)

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

		err = c.habitsModel.UpdateHandler(habit, id)

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

		result, err := c.habitsView.UpdateHandler(habit)

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
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Update - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Update - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Update - Error writing response: %s", err))
	}
}

func (c *HabitsController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.opsChan <- func() {
		id := r.URL.Query().Get("id")

		c.logger.InfoLog(fmt.Sprintf("habitsController.Delete - id=%s", id))

		if len(id) == 0 {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Delete - id query param is missing"),
			}
			return
		}

		err := c.habitsModel.DeleteHandler(id)

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

		result, err := c.habitsView.DeleteHandler()

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
		return
	}

	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	c.logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		c.logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}
