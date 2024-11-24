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
	CreateHabitsHandler(habit data.NewHabit) error
	RetrieveHabitsHandler(id string) (data.Habit, error)
	RetrieveAllHabitsHandler() ([]data.Habit, error)
	UpdateHabitsHandler(habit data.Habit, id string) error
	DeleteHabitsHandler(id string) error
}

func NewHabitsModel(logger logger.ILogger, db db.IDB) *HabitsModel {
	return &HabitsModel{
		logger: logger,
		db:     db,
	}
}

func (m *HabitsModel) CreateHabitsHandler(habit data.NewHabit) error {
	m.logger.InfoLog("habitsModel.Create")

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - err=%s", err))
		return err
	}

	err := m.db.CreateHabitsHandler(habit)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - db.Create - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) RetrieveHabitsHandler(id string) (data.Habit, error) {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Retrieve - id=%s", id))
	habit := data.Habit{}

	result, err := m.db.RetrieveHabitsHandler(id)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Retrieve - db.Retrieve - err=%s", err))
		return habit, err
	}

	habit, ok := result.(data.Habit)

	if !ok {
		return habit, fmt.Errorf("habitsModel.Retrieve - habits type is not data.Habit")
	}
	fmt.Printf("HabitsModel.RetrieveHabitsHandler() returning habit id=%s\n", habit.ID)
	return habit, nil
}

func (m *HabitsModel) RetrieveAllHabitsHandler() ([]data.Habit, error) {
	m.logger.InfoLog("habitsModel.RetrieveAll")
	result, err := m.db.RetrieveAllHabitsHandler()

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

func (m *HabitsModel) UpdateHabitsHandler(habit data.Habit, id string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Update - id=%s", id))

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - err=%s", err))
		return err
	}

	if err := m.db.UpdateHabitsHandler(id, habit); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - db.Update - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) DeleteHabitsHandler(id string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Delete - id=%s", id))
	if err := m.db.DeleteHabitsHandler(id); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Delete - db.Delete - err=%s", err))
		return err
	}

	return nil
}
