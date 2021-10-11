package calendar

import (
	"dev11/pkg/model"
	"fmt"
	"sync"
	"time"
)

// Calendar - interface for calendar
type Calendar interface {
	CreateEvent(model.Event)
	UpdateEvent(model.Event)
	DeleteEvent(model.Event)
	FindEvents(string, time.Time) ([]model.Event, error)
}

// Calendar implements calendar
type calendar struct {
	events map[int]model.Event
	lastID int
	sync.RWMutex
}

// New creates calendar object
func New() Calendar {
	return &calendar{
		events: make(map[int]model.Event),
		lastID: 0,
	}
}

// CreateEvent creates event
func (c *calendar) CreateEvent(event model.Event) {
	c.Lock()
	event.ID = c.lastID
	c.events[event.ID] = event
	c.lastID++
	c.Unlock()
}

// DeleteEvent deletes event
func (c *calendar) DeleteEvent(event model.Event) {
	c.Lock()
	delete(c.events, event.ID)
	c.Unlock()
}

// UpdateEvent updates event
func (c *calendar) UpdateEvent(event model.Event) {
	c.Lock()
	c.events[event.ID] = event
	c.Unlock()
}

// FindEvents finds events for day, week, month
func (c *calendar) FindEvents(time string, date time.Time) ([]model.Event, error) {
	result := make([]model.Event, 0)

	switch time {
	case "day":
		dateStart := date.AddDate(0, 0, -1)
		dateStop := date.AddDate(0, 0, 1)
		for _, v := range c.events {
			if v.Dt.After(dateStart) && v.Dt.Before(dateStop) {
				result = append(result, v)
			}
		}
	case "week":
		dateStart := date.AddDate(0, 0, -7)
		dateStop := date.AddDate(0, 0, 7)
		for _, v := range c.events {
			if v.Dt.After(dateStart) && v.Dt.Before(dateStop) {
				result = append(result, v)
			}
		}
	case "month":
		// fmt.Println("here")
		fmt.Println(c.events)
		dateStart := date.AddDate(0, -1, 0)
		dateStop := date.AddDate(0, 1, 0)
		for _, v := range c.events {
			// fmt.Println("here now")
			if v.Dt.After(dateStart) && v.Dt.Before(dateStop) {
				result = append(result, v)
				// fmt.Println("found event!")
			}
		}
	}
	return result, nil
}
