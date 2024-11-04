package view

import (
	"bytes"
	"dohabits/data"
	"dohabits/logger"
	"encoding/json"
	"testing"
)

func Test_Create(t *testing.T) {
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
			got, err := v.Create(newHabit)

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
			got, err := v.Retrieve(habit)

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
			got, err := v.RetrieveAll(habits)

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
			name: "Test Successful Update",
			want: marshalledHabit,
		},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			got, err := v.Update(habit)

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
			got, err := v.Delete()

			if err != nil {
				t.Errorf("Fail err: %s", err)
			}

			if !bytes.Equal(val.want, got) {
				t.Errorf("Fail want doesn't match got")
			}
		})
	}
}
