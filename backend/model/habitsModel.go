package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/helper"
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
	m.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("habitsModel.Create = userEmailAddress=%s", userEmailAddress))

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("habitsModel.Create - err=%s", err))
		return err
	}

	userDetails, err := m.db.GetUserDetails(&data.UserAuth{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(*data.UserData)

	if !ok {
		return fmt.Errorf(helper.GetFunctionName(), "data.UserData is invalid")
	}

	err = m.db.CreateHabitsHandler(currentUserData.UserID, habit)

	if err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) RetrieveHabitsHandler(userEmailAddress, habitId string) (data.Habit, error) {
	m.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))
	habit := data.Habit{}

	userDetails, err := m.db.GetUserDetails(&data.UserAuth{EmailAddress: userEmailAddress})

	if err != nil {
		return habit, err
	}

	currentUserData, ok := userDetails.(*data.UserData)

	if !ok {
		return data.Habit{}, fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	result, err := m.db.RetrieveHabitsHandler(currentUserData.UserID, habitId)

	if err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return habit, err
	}

	habit, ok = result.(data.Habit)

	if !ok {
		return habit, fmt.Errorf("%s - habits type is not data.Habit", helper.GetFunctionName())
	}

	return habit, nil
}

func (m *HabitsModel) RetrieveAllHabitsHandler(userEmailAddress string) ([]data.Habit, error) {
	m.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userEmailAddress=%s", userEmailAddress))

	userDetails, err := m.db.GetUserDetails(&data.UserAuth{EmailAddress: userEmailAddress})

	if err != nil {
		return nil, err
	}

	currentUserData, ok := userDetails.(*data.UserData)

	if !ok {
		return nil, fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	result, err := m.db.RetrieveAllHabitsHandler(currentUserData.UserID)

	if err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return nil, err
	}

	habits, ok := result.([]data.Habit)

	if !ok {
		return nil, fmt.Errorf("%s - habits type is not data.Habit", helper.GetFunctionName())
	}

	return habits, nil
}

func (m *HabitsModel) UpdateHabitsHandler(userEmailAddress string, habit data.Habit, habitId string) error {
	m.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))

	userDetails, err := m.db.GetUserDetails(&data.UserAuth{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(*data.UserData)

	if !ok {
		return fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	if err := validation.ValidateHabit(habit, m.logger); err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return err
	}

	if err := m.db.UpdateHabitsHandler(currentUserData.UserID, habitId, habit); err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return err
	}

	return nil
}

func (m *HabitsModel) DeleteHabitsHandler(userEmailAddress, habitId string) error {
	m.logger.InfoLog(helper.GetFunctionName(), fmt.Sprintf("userEmailAddress=%s, habitId=%s", userEmailAddress, habitId))

	userDetails, err := m.db.GetUserDetails(&data.UserAuth{EmailAddress: userEmailAddress})

	if err != nil {
		return err
	}

	currentUserData, ok := userDetails.(*data.UserData)

	if !ok {
		return fmt.Errorf("%s - data.UserData is invalid", helper.GetFunctionName())
	}

	if err := m.db.DeleteHabitsHandler(currentUserData.UserID, habitId); err != nil {
		m.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("err=%s", err))
		return err
	}

	return nil
}
