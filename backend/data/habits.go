package data

import "time"

type Habit struct {
	HabitID         string    `json:"habitId" bson:"_id"`
	UserID          string    `json:"userId" bson:"userId"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
	Name            string    `json:"name" bson:"name"`
	Days            int       `json:"days" bson:"days"`
	DaysTarget      int       `json:"daysTarget" bson:"daysTarget"`
	CompletionDates []string  `json:"completionDates" bson:"completionDates"`
}

type NewHabit struct {
	Name       string `json:"name" bson:"name"`
	Days       int    `json:"days" bson:"days"`
	DaysTarget int    `json:"daysTarget" bson:"daysTarget"`
}

type UpdateHabit struct {
	Name            *string   `json:"name" bson:"name"`
	Days            *int      `json:"days" bson:"days"`
	DaysTarget      *int      `json:"daysTarget" bson:"daysTarget"`
	CompletionDates *[]string `json:"completionDates" bson:"completionDates"`
}
