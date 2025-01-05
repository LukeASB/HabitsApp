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
	CreateHabitsHandler(userEmailAddress string, habit data.NewHabit) error
	RetrieveHabitsHandler(userEmailAddress, habitId string) (data.Habit, error)
	RetrieveAllHabitsHandler(userEmailAddress string) ([]data.Habit, error)
	UpdateHabitsHandler(userEmailAddress string, habit data.Habit, habitId string) error
	DeleteHabitsHandler(userEmailAddress, habitId string) error
}

func NewHabitsModel(logger logger.ILogger, db db.IDB) *HabitsModel {
	return &HabitsModel{
		logger: logger,
		db:     db,
	}
}

func (m *HabitsModel) CreateHabitsHandler(userEmailAddress string, habit data.NewHabit) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Create = userEmailAddress=%s", userEmailAddress))

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - err=%s", err))
		return err
	}

	userDetails, err := m.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return fmt.Errorf("authModel.RegisterUserHandler - data.UserData is invalid")
	}

	err = m.db.CreateHabitsHandler(currentUserData.UserID, habit)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Create - db.Create - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) RetrieveHabitsHandler(userEmailAddress, habitId string) (data.Habit, error) {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Retrieve - userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))
	habit := data.Habit{}

	userDetails, err := m.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userEmailAddress})

	if err != nil {
		return habit, err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return data.Habit{}, fmt.Errorf("authModel.RegisterUserHandler - data.UserData is invalid")
	}

	result, err := m.db.RetrieveHabitsHandler(currentUserData.UserID, habitId)

	if err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Retrieve - db.Retrieve - err=%s", err))
		return habit, err
	}

	habit, ok = result.(data.Habit)

	if !ok {
		return habit, fmt.Errorf("habitsModel.Retrieve - habits type is not data.Habit")
	}
	fmt.Printf("HabitsModel.RetrieveHabitsHandler() returning habit id=%s\n", habit.HabitID)
	return habit, nil
}

func (m *HabitsModel) RetrieveAllHabitsHandler(userEmailAddress string) ([]data.Habit, error) {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.RetrieveAll - userEmailAddress=%s", userEmailAddress))

	userDetails, err := m.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userEmailAddress})

	if err != nil {
		return nil, err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return nil, fmt.Errorf("authModel.RegisterUserHandler - data.UserData is invalid")
	}

	result, err := m.db.RetrieveAllHabitsHandler(currentUserData.UserID)

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

func (m *HabitsModel) UpdateHabitsHandler(userEmailAddress string, habit data.Habit, habitId string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Update - userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))

	userDetails, err := m.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return fmt.Errorf("authModel.RegisterUserHandler - data.UserData is invalid")
	}

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - err=%s", err))
		return err
	}

	if err := m.db.UpdateHabitsHandler(currentUserData.UserID, habitId, habit); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Update - db.Update - err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) DeleteHabitsHandler(userEmailAddress, habitId string) error {
	m.logger.InfoLog(fmt.Sprintf("habitsModel.Delete - userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))

	userDetails, err := m.db.GetUserDetails(&data.RegisterUserRequest{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(data.UserData)

	if !ok {
		return fmt.Errorf("authModel.RegisterUserHandler - data.UserData is invalid")
	}

	if err := m.db.DeleteHabitsHandler(currentUserData.UserID, habitId); err != nil {
		m.logger.ErrorLog(fmt.Sprintf("habitsModel.Delete - db.Delete - err=%s", err))
		return err
	}

	return nil
}
