package models

import "time"

type Profile struct {
	UserID         string    `json:"user_id"`
	DOB            time.Time `json:"dob"`
	LifeExpectancy int16     `json:"life_expectancy"`
	WeekStart      string    `json:"week_start"`
	Theme          any       `json:"theme,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
