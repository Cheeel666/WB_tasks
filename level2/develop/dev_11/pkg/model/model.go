package model

import "time"

// Event implements event
type Event struct {
	ID    int       `json:"-"`
	Name  string    `json:"name"`
	Event string    `json:"event"`
	Dt    time.Time `json:"date"`
}
