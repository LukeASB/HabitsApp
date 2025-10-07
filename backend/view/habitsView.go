package view

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"encoding/json"
	"fmt"
)

type HabitsView struct {
	logger logger.ILogger
}

type IHabitsView interface {
	CreateHabitsHandler(newHabit *data.NewHabitResponse) ([]byte, error)
	RetrieveHabitsHandler(habit data.Habit) ([]byte, error)
	RetrieveAllHabitsHandler(habits []data.Habit) ([]byte, error)
	UpdateHabitsHandler(habit data.Habit) ([]byte, error)
	UpdateAllHabitsHandler(habit *[]data.Habit) ([]byte, error)
	DeleteHabitsHandler() ([]byte, error)
}

func NewHabitsView(logger logger.ILogger) *HabitsView {
	return &HabitsView{
		logger: logger,
	}
}

func (v *HabitsView) CreateHabitsHandler(newHabit *data.NewHabitResponse) ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")

	newHabitData := data.NewHabitResponse{
		HabitID:         newHabit.HabitID,
		Name:            newHabit.Name,
		DaysTarget:      newHabit.DaysTarget,
		CompletionDates: []string{},
	}

	result, err := json.Marshal(newHabitData)

	if err != nil {
		v.logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("Error encoding to JSON - err=%s", err))
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveHabitsHandler(habit data.Habit) ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")
	result, err := json.Marshal(habit)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) RetrieveAllHabitsHandler(habits []data.Habit) ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")

	if len(habits) == 0 {
		return []byte("[]"), nil
	}

	result, err := json.Marshal(habits)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) UpdateHabitsHandler(habit data.Habit) ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")

	updatedHabit := data.UpdateHabit{
		Name:            &habit.Name,
		Days:            &habit.Days,
		DaysTarget:      &habit.DaysTarget,
		CompletionDates: &habit.CompletionDates,
	}
	result, err := json.Marshal(updatedHabit)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) UpdateAllHabitsHandler(habit *[]data.Habit) ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")

	var updatedHabit []data.UpdateHabit

	for _, val := range *habit {
		name := val.Name
		days := val.Days
		daysTarget := val.DaysTarget
		completionDates := val.CompletionDates

		updatedHabit = append(updatedHabit, data.UpdateHabit{
			HabitID:         val.HabitID,
			Name:            &name,
			Days:            &days,
			DaysTarget:      &daysTarget,
			CompletionDates: &completionDates,
		})
	}
	result, err := json.Marshal(updatedHabit)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (v *HabitsView) DeleteHabitsHandler() ([]byte, error) {
	v.logger.InfoLog(helper.GetFunctionName(), "")
	response := map[string]bool{"success": true}

	jsonRes, err := json.Marshal(response)

	if err != nil {
		return nil, fmt.Errorf("%s - failed to encode response: %v", helper.GetFunctionName(), err)
	}

	return jsonRes, nil
}
