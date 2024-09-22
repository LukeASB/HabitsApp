package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"fmt"
)

type HabitsModel struct{}

type IHabitsModel interface {
	Create(habit data.NewHabit, db db.IDB, logger logger.ILogger) error
	Retrieve(id string, db db.IDB, logger logger.ILogger) (data.Habit, error)
	RetrieveAll(db db.IDB, logger logger.ILogger) ([]data.Habit, error)
	Update(habit data.Habit, id string, db db.IDB, logger logger.ILogger) error
	Delete(id string, db db.IDB, logger logger.ILogger) error
}

/*
 Need validate the habits in model layer before they hit the DB. Someone shouldn't be able to call the endpoint to:
 - Create Habits with random symbols/text
 - Enter negative days
*/

func (m *HabitsModel) Create(habit data.NewHabit, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog("habitsModel.Create")
	err := db.Create(logger, habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Create - db.Create - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) Retrieve(id string, db db.IDB, logger logger.ILogger) (data.Habit, error) {
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

func (m *HabitsModel) RetrieveAll(db db.IDB, logger logger.ILogger) ([]data.Habit, error) {
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

func (m *HabitsModel) Update(habit data.Habit, id string, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog(fmt.Sprintf("habitsModel.Update - id=%s", id))
	if err := db.Update(logger, id, habit); err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Update - db.Update - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) Delete(id string, db db.IDB, logger logger.ILogger) error {
	logger.InfoLog(fmt.Sprintf("habitsModel.Delete - id=%s", id))
	if err := db.Delete(logger, id); err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsModel.Delete - db.Delete - err=%s", err))
		return err
	}

	return nil
}
