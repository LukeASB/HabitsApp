package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"dohabits/validation"
	"fmt"
)

type HabitsModel struct {
	logger logger.ILogger
	db     db.IDB
}

type IHabitsModel interface {
	CreateHandler(habit data.NewHabit) error
	RetrieveHandler(id string) (data.Habit, error)
	RetrieveAllHandler() ([]data.Habit, error)
	UpdateHandler(habit data.Habit, id string) error
	DeleteHandler(id string) error
}

func NewHabitsModel(logger logger.ILogger, db db.IDB) *HabitsModel {
	return &HabitsModel{
		logger: logger,
		db:     db,
	}
}

/*
 Need validate the habits in model layer before they hit the DB. Someone shouldn't be able to call the endpoint to:
 - Create Habits with random symbols/text
 - Enter negative days
*/

func (m *HabitsModel) CreateHandler(habit data.NewHabit) error {
	m.logger.InfoLog("habitsModel.Create")

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - err=%s", err))
		return err
	}

	err := m.db.CreateHandler(habit)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - db.Create - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) RetrieveHandler(id string) (data.Habit, error) {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Retrieve - id=%s", id))
	habit := data.Habit{}

	result, err := m.db.RetrieveHandler(id)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Retrieve - db.Retrieve - err=%s", err))
		return habit, err
	}

	habit, ok := result.(data.Habit)

	if !ok {
		return habit, fmt.Errorf("habitsModel.Retrieve - habits type is not data.Habit")
	}
	fmt.Printf("HabitsModel.RetrieveHandler() returning habit id=%s\n", habit.ID)
	return habit, nil
}

func (m *HabitsModel) RetrieveAllHandler() ([]data.Habit, error) {
	m.logger.InfoLog("habitsModel.RetrieveAll")
	result, err := m.db.RetrieveAllHandler()

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.RetrieveAll - db.RetrieveAll - err=%s", err))
		return nil, err
	}

	habits, ok := result.([]data.Habit)

	if !ok {
		return nil, fmt.Errorf("habitsModel.RetrieveAll - habits type is not data.Habit")
	}

	return habits, nil
}

func (m *HabitsModel) UpdateHandler(habit data.Habit, id string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Update - id=%s", id))

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - err=%s", err))
		return err
	}

	if err := m.db.UpdateHandler(id, habit); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - db.Update - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) DeleteHandler(id string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Delete - id=%s", id))
	if err := m.db.DeleteHandler(id); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Delete - db.Delete - err=%s", err))
		return err
	}

	return nil
}
