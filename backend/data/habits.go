package data

import "time"

type Habit struct {
	HabitID         string    `json:"habitId"`
	UserID          string    `json:"userId"`
	CreatedAt       time.Time `json:"createdAt"`
	Name            string    `json:"name"`
	Days            int       `json:"days"`
	DaysTarget      int       `json:"daysTarget"`
	CompletionDates []string  `json:"completionDates"`
}

type NewHabit struct {
	Name       string `json:"name"`
	Days       int    `json:"days"`
	DaysTarget int    `json:"daysTarget"`
}

type UpdateHabit struct {
	Name       *string `json:"name"`
	Days       *int    `json:"days"`
	DaysTarget *int    `json:"daysTarget"`
}
