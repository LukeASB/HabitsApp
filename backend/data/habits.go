package data

import "time"

type Habit struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	Days       int       `json:"days"`
	DaysTarget int       `json:"daysTarget"`
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
