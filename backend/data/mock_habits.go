package data

import (
	"time"
)

var MockHabitJSON = `[
    {
        "ID": "1",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Actually Finish This Project",
        "Days": 30,
        "DaysTarget": 66
    },
    {
        "ID": "2",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Code everyday",
        "Days": 30,
        "DaysTarget": 66
    },
    {
        "ID": "3",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Wake up at 5am everyday",
        "Days": 5,
        "DaysTarget": 365
    },
    {
        "ID": "4",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Cold shower everyday",
        "Days": 25,
        "DaysTarget": 30
    },
    {
        "ID": "5",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Read one book a week",
        "Days": 30,
        "DaysTarget": 30
    },
    {
        "ID": "6",
        "CreatedAt": "2024-09-01T00:00:00Z",
        "Name": "Limit phone screen time to 1 hour a day",
        "Days": 5,
        "DaysTarget": 60
    }
]`

var MockHabit = []Habit{
	{
		ID:         "1",
		CreatedAt:  time.Now(),
		Name:       "Actually Finish This Project",
		Days:       30,
		DaysTarget: 66,
	},
	{
		ID:         "2",
		CreatedAt:  time.Now(),
		Name:       "Code everyday",
		Days:       30,
		DaysTarget: 66,
	},
	{
		ID:         "3",
		CreatedAt:  time.Now(),
		Name:       "Wake up at 5am everyday",
		Days:       5,
		DaysTarget: 365,
	},
	{
		ID:         "4",
		CreatedAt:  time.Now(),
		Name:       "Cold shower everyday",
		Days:       25,
		DaysTarget: 30,
	},
	{
		ID:         "5",
		CreatedAt:  time.Now(),
		Name:       "Read one book a week",
		Days:       30,
		DaysTarget: 30,
	},
	{
		ID:         "6",
		CreatedAt:  time.Now(),
		Name:       "Limit phone screen time to 1 hour a day",
		Days:       5,
		DaysTarget: 60,
	},
}
