package validation

import (
	"dohabits/data"
	"dohabits/logger"
	"testing"
)

func Test_ValidateHabit(t *testing.T) {
	logger := &logger.Logger{}

	testCases := []struct {
		name    string
		habit   data.Habit
		wantErr bool
	}{
		{
			name:    "Create a success valid habit",
			habit:   data.Habit{Name: "Test Habit", Days: 66, DaysTarget: 88},
			wantErr: false,
		},
		{
			name:    "Invalid Habit Name",
			habit:   data.Habit{Name: "%Habit''111", Days: 1000, DaysTarget: 1000},
			wantErr: true,
		},
		{
			name:    "Invalid Habit Days",
			habit:   data.Habit{Name: "Habit Name", Days: 99999, DaysTarget: 99999},
			wantErr: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got := ValidateHabit(val.habit, logger)

			if val.wantErr != (got != nil) {
				t.Errorf("Fail want not equal to got. wantErr: %v, got: %s", val.wantErr, got)
			}
		})
	}
}

func Test_ValidateHabitName(t *testing.T) {
	testCases := []struct {
		name    string
		habit   data.Habit
		wantErr bool
	}{
		{
			name:    "Valid Habit Name",
			habit:   data.Habit{Name: "Good Habit Name", Days: 1, DaysTarget: 30},
			wantErr: false,
		},
		{
			name:    "ValidHabitName",
			habit:   data.Habit{Name: "GoodHabitName", Days: 1, DaysTarget: 30},
			wantErr: false,
		},
		{
			name:    "Valid Habit Example - wake up at 5:00am",
			habit:   data.Habit{Name: "Wake up at 5:00 everyday", Days: 1, DaysTarget: 30},
			wantErr: false,
		},
		{
			name:    "Invalid Habit Name - no name supplied",
			habit:   data.Habit{Name: "", Days: 1, DaysTarget: 30},
			wantErr: true,
		},
		{
			name:    "Invalid Habit Name - Longer than 255 characters limits",
			habit:   data.Habit{Name: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWX", Days: 1, DaysTarget: 30},
			wantErr: true,
		},
		{
			name:    "Invalid Habit Name - symbols",
			habit:   data.Habit{Name: "Bad H@bit N@me!", Days: 1, DaysTarget: 30},
			wantErr: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got := validateHabitName(val.habit.Name)

			if val.wantErr != (got != nil) {
				t.Errorf("Fail want not equal to got. wantErr: %v, got: %s", val.wantErr, got)
			}
		})
	}
}

func Test_ValidateHabitDaysTarget(t *testing.T) {
	testCases := []struct {
		name  string
		habit data.Habit

		wantErr bool
	}{
		{
			name:    "Valid Habit DaysTarget",
			habit:   data.Habit{Name: "Good Habit Name", Days: 1, DaysTarget: 30},
			wantErr: false,
		},
		{
			name:    "Invalid Habit Days - Exceeds Max Days Target",
			habit:   data.Habit{Name: "Good Habit Name", Days: 1, DaysTarget: 99999},
			wantErr: true,
		},
		{
			name:    "Invalid Habit Days - Negative DaysTarget",
			habit:   data.Habit{Name: "Good Habit Name", Days: 1, DaysTarget: -1},
			wantErr: true,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got := validateHabitDaysTarget(val.habit.DaysTarget)

			if val.wantErr != (got != nil) {
				t.Errorf("Fail want not equal to got. wantErr: %v, got: %s", val.wantErr, got)
			}
		})
	}
}
