package view

import (
	"bytes"
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"testing"
)

func Test_CreateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	v := NewHabitsView(logger)

	newHabit := data.NewHabit{
		Name:       "Test Create Habit",
		Days:       11,
		DaysTarget: 30,
	}

	marshalledNewHabit, err := json.Marshal(newHabit)

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test Successful Create",
			want: marshalledNewHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.CreateHabitsHandler(newHabit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_RetrieveHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	v := NewHabitsView(logger)

	habit := data.MockHabit[0]

	marshalledHabit, err := json.Marshal(habit)

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test Successful Create",
			want: marshalledHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.RetrieveHabitsHandler(habit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_RetrieveAllHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	v := NewHabitsView(logger)

	habits := data.MockHabit

	marshalledHabits, err := json.Marshal(habits)

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test Successful Retrieve All",
			want: marshalledHabits,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.RetrieveAllHabitsHandler(habits)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_UpdateHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	v := NewHabitsView(logger)

	habit := data.Habit{
		Name:            "Test Update Habit",
		Days:            30,
		DaysTarget:      50,
		CompletionDates: append(data.MockHabit[0].CompletionDates, []string{"2021-09-01", "2021-09-02", "2021-09-03"}...),
	}

	marshalledUpdatedHabit, err := json.Marshal(data.UpdateHabit{
		Name:            &habit.Name,
		Days:            &habit.Days,
		DaysTarget:      &habit.DaysTarget,
		CompletionDates: &habit.CompletionDates,
	})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test Successful Update",
			want: marshalledUpdatedHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.UpdateHabitsHandler(habit)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_DeleteHabitsHandler(t *testing.T) {
	logger := &logger.Logger{}
	v := NewHabitsView(logger)

	want, err := json.Marshal(map[string]bool{"success": true})

	if err != nil {
		t.Errorf("Fail err: %s", err)
	}

	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test Successful Delete",
			want: want,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.DeleteHabitsHandler()

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}
