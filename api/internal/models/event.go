package models

import "time"

type Event struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	DateStart   time.Time  `json:"date_start"`
	DateEnd     *time.Time `json:"date_end,omitempty"`
	Tag         *string    `json:"tag,omitempty"`
	Color       *string    `json:"color,omitempty"`
	Visibility  string     `json:"visibility"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
