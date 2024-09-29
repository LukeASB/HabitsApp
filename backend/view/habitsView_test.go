package view

import (
	"bytes"
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"testing"
)

func Test_Create(t *testing.T) {
	v := &HabitsView{}
	logger := &logger.Logger{}

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
			got, err := v.Create(newHabit, logger)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_Retrieve(t *testing.T) {
	v := &HabitsView{}
	logger := &logger.Logger{}

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
			got, err := v.Retrieve(habit, logger)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_RetrieveAll(t *testing.T) {
	v := &HabitsView{}
	logger := &logger.Logger{}

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
			got, err := v.RetrieveAll(habits, logger)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_Update(t *testing.T) {
	v := &HabitsView{}
	logger := &logger.Logger{}

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
			name: "Test Successful Update",
			want: marshalledHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.Update(habit, logger)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	v := &HabitsView{}
	logger := &logger.Logger{}

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
			got, err := v.Delete(logger)

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}
