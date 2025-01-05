package view

import (
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"fmt"
)

type HabitsView struct {
	logger logger.ILogger
}

type IHabitsView interface {
	CreateHabitsHandler(newHabit data.NewHabit) ([]byte, error)
	RetrieveHabitsHandler(habit data.Habit) ([]byte, error)
	RetrieveAllHabitsHandler(habits []data.Habit) ([]byte, error)
	UpdateHabitsHandler(habit data.Habit) ([]byte, error)
	DeleteHabitsHandler() ([]byte, error)
}

func NewHabitsView(logger logger.ILogger) *HabitsView {
	return &HabitsView{
		logger: logger,
	}
}

func (v *HabitsView) CreateHabitsHandler(newHabit data.NewHabit) ([]byte, error) {
	v.logger.InfoLog("habitsView.Create")
	result, err := json.Marshal(newHabit)

	if err != nil {
		v.logger.ErrorLog(fmt.Sprintf("habitsView.Create - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveHabitsHandler(habit data.Habit) ([]byte, error) {
	v.logger.InfoLog("habitsView.Retrieve")
	result, err := json.Marshal(habit)

	if err != nil {
		v.logger.ErrorLog(fmt.Sprintf("habitsView.Retrieve - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveAllHabitsHandler(habits []data.Habit) ([]byte, error) {
	v.logger.InfoLog("habitsView.RetrieveAll")
	result, err := json.Marshal(habits)

	if err != nil {
		v.logger.ErrorLog(fmt.Sprintf("habitsView.RetrieveAll - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) UpdateHabitsHandler(habit data.Habit) ([]byte, error) {
	v.logger.InfoLog("habitsView.Update")

	updatedHabit := data.UpdateHabit{
		Name:            &habit.Name,
		Days:            &habit.Days,
		DaysTarget:      &habit.DaysTarget,
		CompletionDates: &habit.CompletionDates,
	}
	result, err := json.Marshal(updatedHabit)

	if err != nil {
		v.logger.ErrorLog(fmt.Sprintf("habitsView.Update - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) DeleteHabitsHandler() ([]byte, error) {
	v.logger.InfoLog("habitsView.Delete")
	response := map[string]bool{"success": true}

	jsonRes, err := json.Marshal(response)

	if err != nil {
		v.logger.ErrorLog(fmt.Sprintf("habitsView.Delete - Error encoding to JSON - err=%s", err))
		return nil, fmt.Errorf("failed to encode response: %v", err)
	}

	return jsonRes, nil
}
