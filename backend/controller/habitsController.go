package controller

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/model"
	"dohabits/view"
	"encoding/json"
	"fmt"
	"net/http"
)

type HabitsController struct {
	OpsChan chan func()
}

type IHabitsController interface {
	Create(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger)
	Retrieve(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger)
	RetrieveAll(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger)
	Update(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger)
	Delete(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger)
}

// Initialize the processOperations Goroutine
func NewHabitsController(logger logger.ILogger) *HabitsController {
	logger.InfoLog("habitsController.NewHabitsController")

	habitsController := &HabitsController{
		OpsChan: make(chan func(), 1),
	}

	go habitsController.manageOps(logger) // Run the Goroutine to handle operations.

	return habitsController
}

func (c *HabitsController) manageOps(logger logger.ILogger) {
	logger.InfoLog("habitsController.manageOps")
	// Wait and execute any function that get sent to OpsChan
	for op := range c.OpsChan {
		logger.DebugLog("habitsController.manageOps - exec func passed to channel")
		op() // Execute the function passed to the channel
	}
}

func (c *HabitsController) Create(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.OpsChan <- func() {
		logger.InfoLog("habitsController.Create")

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

		if err := m.Create(newHabit, db, logger); err != nil {
			resultChan <- struct {
				data []byte
				err  error
			}{
				data: []byte(""),
				err:  fmt.Errorf("habitsController.Create - model.Create - err=%s", err),
			}
			return
		}

		result, err := v.Create(newHabit, logger)

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
		logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}

func (c *HabitsController) Retrieve(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.OpsChan <- func() {
		id := r.URL.Query().Get("id")

		logger.InfoLog(fmt.Sprintf("habitsController.Retrieve - id=%s", id))

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

		habit, err := m.Retrieve(id, db, logger)

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

		result, err := v.Retrieve(habit, logger)

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
		logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	logger.DebugLog(fmt.Sprintf("habitsController.Retrieve - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsController.Retrieve - Error writing response: %s", err))
	}
}

func (c *HabitsController) RetrieveAll(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.OpsChan <- func() {
		logger.InfoLog("habitsController.RetrieveAll")

		habits, err := m.RetrieveAll(db, logger)

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

		result, err := v.RetrieveAll(habits, logger)

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
		logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	logger.DebugLog(fmt.Sprintf("habitsController.RetrieveAll - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsController.RetrieveAll - Error writing response: %s", err))
	}
}

func (c *HabitsController) Update(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.OpsChan <- func() {
		id := r.URL.Query().Get("id")

		logger.InfoLog(fmt.Sprintf("habitsController.Update - id=%s", id))

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

		habit, err := m.Retrieve(id, db, logger)

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

		err = m.Update(habit, id, db, logger)

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

		result, err := v.Update(habit, logger)

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
		logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logger.DebugLog(fmt.Sprintf("habitsController.Update - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	logger.DebugLog(fmt.Sprintf("habitsController.Update - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsController.Update - Error writing response: %s", err))
	}
}

func (c *HabitsController) Delete(w http.ResponseWriter, r *http.Request, m model.IHabitsModel, v view.IHabitsView, db db.IDB, logger logger.ILogger) {
	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	resultChan := make(chan struct {
		data []byte
		err  error
	}, 1)

	c.OpsChan <- func() {
		id := r.URL.Query().Get("id")

		logger.InfoLog(fmt.Sprintf("habitsController.Delete - id=%s", id))

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

		err := m.Delete(id, db, logger)

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

		result, err := v.Delete(logger)

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
		logger.ErrorLog(fmt.Sprintf("%s", res.err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logger.DebugLog(fmt.Sprintf("habitsController.Delete - Writing response: %s", res.data))
	numOfBytes, err := w.Write([]byte(res.data))
	logger.DebugLog(fmt.Sprintf("habitsController.Delete - w.Write wrote %d bytes", numOfBytes))
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsController.Delete - Error writing response: %s", err))
	}
}
