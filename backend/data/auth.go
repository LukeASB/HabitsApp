package data

import "time"

type Login struct {
	Success     bool      `json:"succcess"`
	AccessToken string    `json:"accessToken"`
	LoggedInAt  time.Time `json:"LoggedInAt"`
}
