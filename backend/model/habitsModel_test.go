package model

import (
	"dohabits/data"
	"dohabits/db"
	"dohabits/logger"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	db := &db.MyMockDB{}
	model := &HabitsModel{}
	logger := &logger.Logger{}

	testCases := []struct {
		name     string
		newHabit data.NewHabit
		want     error
	}{
		{
			name:     "Successfully Create Habit",
			newHabit: data.NewHabit{},
			want:     nil,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			if err := model.Create(val.newHabit, db, logger); err != nil {
				t.Errorf("TestCreate Failed - err=%s", err)
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	db := &db.MyMockDB{}
	model := &HabitsModel{}
	logger := &logger.Logger{}

	testCases := []struct {
		name string
		id   string
		want data.Habit
	}{
		{
			name: "Get Habit Successfully",
			id:   data.MockHabit[0].ID,
			want: data.MockHabit[0],
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			habit, err := model.Retrieve(val.id, db, logger)

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

func TestRetrieveAll(t *testing.T) {
	db := &db.MyMockDB{}
	model := &HabitsModel{}
	logger := &logger.Logger{}

	testCases := []struct {
		name string
		want []data.Habit
	}{
		{
			name: "Get All Habits Successfully",
			want: data.MockHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			habits, err := model.RetrieveAll(db, logger)

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

func TestUpdate(t *testing.T) {
	db := &db.MyMockDB{}
	model := &HabitsModel{}
	logger := &logger.Logger{}

	testCases := []struct {
		name        string
		updateHabit data.Habit
		id          string
	}{
		{
			name:        "Update Habit Successfully",
			updateHabit: data.Habit{Name: "Pray everday", Days: 12, DaysTarget: 30},
			id:          "1",
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			if err := model.Update(val.updateHabit, val.id, db, logger); err != nil {
				t.Errorf("TestUpdate Failed - err=%s", err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	db := &db.MyMockDB{}
	model := &HabitsModel{}
	logger := &logger.Logger{}

	testCases := []struct {
		name    string
		id      string
		wantErr bool
		got     bool
	}{
		{
			name:    "Delete Habit Successfully",
			id:      "1",
			wantErr: false,
		},
		{
			name:    "Delete Habit Failed",
			id:      "9",
			wantErr: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			err := model.Delete(val.id, db, logger)

			if err != nil {
				val.got = true
			}

			if val.wantErr != val.got {
				t.Errorf("TestDelete Failed - err=%s", err)
			}
		})
	}
}
