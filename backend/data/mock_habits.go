package data

import (
	"time"
)

var MockHabit = []Habit{
	{
		HabitID:         "1",
		UserID:          "1",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Actually Finish This Project",
		Days:            30,
		DaysTarget:      66,
		CompletionDates: []string{"2024-12-20", "2024-12-02", "2024-12-03"},
	},
	{
		HabitID:         "2",
		UserID:          "1",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Code everyday",
		Days:            30,
		DaysTarget:      66,
		CompletionDates: []string{"2024-12-20", "2024-12-02", "2024-12-11"},
	},
	{
		HabitID:         "3",
		UserID:          "2",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Wake up at 5am everyday",
		Days:            5,
		DaysTarget:      365,
		CompletionDates: []string{"2024-12-20", "2024-12-11", "2024-12-11"},
	},
	{
		HabitID:         "4",
		UserID:          "3",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Cold shower everyday",
		Days:            25,
		DaysTarget:      30,
		CompletionDates: []string{"2024-12-20", "2024-12-12", "2024-12-11"},
	},
	{
		HabitID:         "5",
		UserID:          "4",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Read one book a week",
		Days:            30,
		DaysTarget:      30,
		CompletionDates: []string{"2024-12-20", "2024-12-02", "2024-12-11"},
	},
	{
		HabitID:         "6",
		UserID:          "1",
		CreatedAt:       time.Date(2024, time.September, 21, 10, 30, 0, 0, time.UTC),
		Name:            "Limit phone screen time to 1 hour a day",
		Days:            5,
		DaysTarget:      60,
		CompletionDates: []string{"2024-12-20", "2024-12-02", "2024-12-11"},
	},
}
