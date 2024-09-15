package view

import (
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

type HabitsView struct{}

type IHabitsView interface {
	Create(w http.ResponseWriter, r *http.Request, newHabit data.NewHabit, logger logger.ILogger) ([]byte, error)
	Retrieve(w http.ResponseWriter, r *http.Request, habit data.Habit, logger logger.ILogger) ([]byte, error)
	RetrieveAll(w http.ResponseWriter, r *http.Request, habits []data.Habit, logger logger.ILogger) ([]byte, error)
	Update(w http.ResponseWriter, r *http.Request, habit data.Habit, logger logger.ILogger) ([]byte, error)
	Delete(w http.ResponseWriter, r *http.Request, logger logger.ILogger) ([]byte, error)
}

func (v *HabitsView) Create(w http.ResponseWriter, r *http.Request, newHabit data.NewHabit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Create")
	result, err := json.Marshal(newHabit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Create - Error encoding to JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Retrieve(w http.ResponseWriter, r *http.Request, habit data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Retrieve")
	result, err := json.Marshal(habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Retrieve - Error encoding to JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveAll(w http.ResponseWriter, r *http.Request, habits []data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.RetrieveAll")
	result, err := json.Marshal(habits)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.RetrieveAll - Error encoding to JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Update(w http.ResponseWriter, r *http.Request, habit data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Update")
	result, err := json.Marshal(habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Update - Error encoding to JSON - err=%s", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Delete(w http.ResponseWriter, r *http.Request, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Delete")
	response := map[string]bool{"success": true}

	w.Header().Set("Content-Type", "application/json")

	jsonRes, err := json.Marshal(response)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Delete - Error encoding to JSON - err=%s", err))
		return nil, fmt.Errorf("failed to encode response: %v", err)
	}

	return jsonRes, nil
}
