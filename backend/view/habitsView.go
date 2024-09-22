package view

import (
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"fmt"
)

type HabitsView struct{}

type IHabitsView interface {
	Create(newHabit data.NewHabit, logger logger.ILogger) ([]byte, error)
	Retrieve(habit data.Habit, logger logger.ILogger) ([]byte, error)
	RetrieveAll(habits []data.Habit, logger logger.ILogger) ([]byte, error)
	Update(habit data.Habit, logger logger.ILogger) ([]byte, error)
	Delete(logger logger.ILogger) ([]byte, error)
}

func (v *HabitsView) Create(newHabit data.NewHabit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Create")
	result, err := json.Marshal(newHabit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Create - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Retrieve(habit data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Retrieve")
	result, err := json.Marshal(habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Retrieve - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveAll(habits []data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.RetrieveAll")
	result, err := json.Marshal(habits)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.RetrieveAll - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Update(habit data.Habit, logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Update")
	result, err := json.Marshal(habit)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Update - Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) Delete(logger logger.ILogger) ([]byte, error) {
	logger.InfoLog("habitsView.Delete")
	response := map[string]bool{"success": true}

	jsonRes, err := json.Marshal(response)

	if err != nil {
		logger.ErrorLog(fmt.Sprintf("habitsView.Delete - Error encoding to JSON - err=%s", err))
		return nil, fmt.Errorf("failed to encode response: %v", err)
	}

	return jsonRes, nil
}
