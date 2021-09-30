package model

import "time"

// Calendar implements calendar
type Calendar struct {
	// Тут должна быть модель данных календаря
}

// CreateEvent represents BI - тут можно короче сделать любую бизнес логику, но, как понял по заданию
// заглушек будет достаточно, так как нет бд и не известна модель данных
func (c *Calendar) CreateEvent(id int, name, event string, dt time.Time) error {
	return nil
}

// DeleteEvent represents BI - тут можно короче сделать любую бизнес логику, но, как понял по заданию
// заглушек будет достаточно, так как нет бд и не известна модель данных
func (c *Calendar) DeleteEvent(eventID int) error {
	return nil
}

// UpdateEvent represents BI - тут можно короче сделать любую бизнес логику, но, как понял по заданию
// заглушек будет достаточно, так как нет бд и не известна модель данных
func (c *Calendar) UpdateEvent(eventID int, event string) error {
	return nil
}

// FindEvents finds events for day, week, month
func (c *Calendar) FindEvents(time string) (map[string][]string, error) {
	var res map[string][]string

	switch time {
	case "day":
		return res, nil // Тут уже тоже какая-то внутренняя логика
	case "week":
		return res, nil
	case "month":
		return res, nil
	}
	return res, nil
}
