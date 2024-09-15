package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"fmt"
	"net/http"
)

type HabitsModel struct{}

type IHabitsModel interface {
	Create(w http.ResponseWriter, r *http.Request, habit data.NewHabit, db db.IDB, logger logger.ILogger) error
	Retrieve(w http.ResponseWriter, r *http.Request, id string, db db.IDB, logger logger.ILogger) (data.Habit, error)
	RetrieveAll(w http.ResponseWriter, r *http.Request, db db.IDB, logger logger.ILogger) ([]data.Habit, error)
	Update(w http.ResponseWriter, r *http.Request, habit data.Habit, id string, db db.IDB, logger logger.ILogger) error
	Delete(w http.ResponseWriter, r *http.Request, id string, db db.IDB, logger logger.ILogger) error
}

func (m *HabitsModel) Create(w http.ResponseWriter, r *http.Request, habit data.NewHabit, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog("habitsModel.Create")
	err := db.Create(logger, habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Create - db.Create - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) Retrieve(w http.ResponseWriter, r *http.Request, id string, db db.IDB, logger logger.ILogger) (data.Habit, error) {
	logger.InfoLog(fmt.Sprintf("habitsModel.Retrieve - id=%s", id))
	habit := data.Habit{}

	result, err := db.Retrieve(logger, id)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Retrieve - db.Retrieve - err=%s", err))
		return habit, err
	}

	habit, ok := result.(data.Habit)

	if !ok {
		return habit, fmt.Errorf("habitsModel.Retrieve - habits type is not data.Habit")
	}
	fmt.Printf("HabitsModel.Retrieve() returning habit id=%s\n", habit.ID)
	return habit, nil
}

func (m *HabitsModel) RetrieveAll(w http.ResponseWriter, r *http.Request, db db.IDB, logger logger.ILogger) ([]data.Habit, error) {
	logger.InfoLog("habitsModel.RetrieveAll")
	result, err := db.RetrieveAll(logger)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.RetrieveAll - db.RetrieveAll - err=%s", err))
		return nil, err
	}

	habits, ok := result.([]data.Habit)

	if !ok {
		return nil, fmt.Errorf("habitsModel.RetrieveAll - habits type is not data.Habit")
	}

	return habits, nil
}

func (m *HabitsModel) Update(w http.ResponseWriter, r *http.Request, habit data.Habit, id string, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog(fmt.Sprintf("habitsModel.Update - id=%s", id))
	if err := db.Update(logger, id, habit); err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Update - db.Update - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) Delete(w http.ResponseWriter, r *http.Request, id string, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog(fmt.Sprintf("habitsModel.Delete - id=%s", id))
	if err := db.Delete(logger, id); err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Delete - db.Delete - err=%s", err))
		return err
	}

	return nil
}
