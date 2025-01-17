package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"reflect"
	"testing"
)

func TestCreateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	model := NewHabitsModel(logger, db)

	testCases := []struct {
		name             string
		userEmailAddress string
		newHabit         data.NewHabit
		want             error
	}{
		{
			name:             "Successfully Create Habit",
			userEmailAddress: "johndoe1@example.com",
			newHabit:         data.NewHabit{Name: "Create Habit Test", Days: 1, DaysTarget: 11},
			want:             nil,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			if err := model.CreateHabitsHandler(val.userEmailAddress, val.newHabit); err != nil {
				t.Errorf("TestCreate Failed - err=%s", err)
			}
		})
	}
}

func TestRetrieveHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	model := NewHabitsModel(logger, db)

	testCases := []struct {
		name             string
		userEmailAddress string
		habitId          string
		want             data.Habit
	}{
		{
			name:             "Get Habit Successfully",
			userEmailAddress: "johndoe1@example.com",
			habitId:          data.MockHabit[0].HabitID,
			want:             data.MockHabit[0],
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			habit, err := model.RetrieveHabitsHandler(val.userEmailAddress, val.habitId)

			if err != nil {
				t.Errorf("TestRetrieve Failed - err=%s", err)
			}

			habitsMatch := reflect.DeepEqual(habit, val.want)

			if habitsMatch == false {
				t.Errorf("TestRetrieve Failed - err=%s", err)
			}
		})
	}
}

func TestRetrieveAllHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	model := NewHabitsModel(logger, db)

	var mockHabitForUserID1 []data.Habit

	for _, habit := range data.MockHabit {
		if habit.UserID == "1" {
			mockHabitForUserID1 = append(mockHabitForUserID1, habit)
		}
	}

	testCases := []struct {
		name             string
		userEmailAddress string
		want             []data.Habit
	}{
		{
			name:             "Get All Habits Successfully",
			userEmailAddress: "johndoe1@example.com",
			want:             mockHabitForUserID1,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			habits, err := model.RetrieveAllHabitsHandler(val.userEmailAddress)

			if err != nil {
				t.Errorf("TestRetrieve Failed - err=%s", err)
			}

			for i, habit := range habits {
				habitsMatch := reflect.DeepEqual(habit, val.want[i])

				if habitsMatch == false {
					t.Errorf("TestRetrieve Failed - err=%s", err)
				}
			}
		})
	}
}

func TestUpdateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	model := NewHabitsModel(logger, db)

	testCases := []struct {
		name             string
		userEmailAddress string
		updateHabit      data.Habit
		habitId          string
	}{
		{
			name:             "Update Habit Successfully",
			userEmailAddress: "johndoe1@example.com",
			updateHabit:      data.Habit{Name: "Pray everday", Days: 12, DaysTarget: 30},
			habitId:          "1",
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			if err := model.UpdateHabitsHandler(val.userEmailAddress, val.updateHabit, val.habitId); err != nil {
				t.Errorf("TestUpdate Failed - err=%s", err)
			}
		})
	}
}

func TestDeleteHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	db := db.NewMockDB(logger)
	model := NewHabitsModel(logger, db)

	testCases := []struct {
		name             string
		userEmailAddress string
		habitId          string
		wantErr          bool
		got              bool
	}{
		{
			name:             "Delete Habit Successfully",
			userEmailAddress: "johndoe1@example.com",
			habitId:          "1",
			wantErr:          false,
		},
		{
			name:             "Delete Habit Failed",
			userEmailAddress: "1",
			habitId:          "9",
			wantErr:          true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			err := model.DeleteHabitsHandler(val.userEmailAddress, val.habitId)

			if err != nil {
				val.got = true
			}

			if val.wantErr != val.got {
				t.Errorf("TestDelete Failed - err=%s", err)
			}
		})
	}
}
