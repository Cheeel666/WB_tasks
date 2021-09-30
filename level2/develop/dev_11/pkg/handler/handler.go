package handler

import (
	"dev11/pkg/middleware"
	"net/http"
)

// Handler type implements handlers
type Handler struct {
}

// InitRoutes initilize routes
func (h *Handler) InitRoutes() {
	// Обьявляем хендлеры, запросы для которых логгируются через middleware
	http.Handle("/create_event", middleware.Logging(http.HandlerFunc(h.CreateEvent)))
	http.Handle("/update_event", middleware.Logging(http.HandlerFunc(h.UpdateEvent)))
	http.Handle("/delete_event", middleware.Logging(http.HandlerFunc(h.DeleteEvent)))

	http.Handle("/events_for_day", middleware.Logging(http.HandlerFunc(h.EventsPerDay)))
	http.Handle("/events_for_week", middleware.Logging(http.HandlerFunc(h.EventsPerWeek)))
	http.Handle("/events_for_month", middleware.Logging(http.HandlerFunc(h.EventsPerMonth)))
}
