package model

import "time"

// Event implements event
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Event  string    `json:"event"`
	Dt     time.Time `json:"date"`
}
