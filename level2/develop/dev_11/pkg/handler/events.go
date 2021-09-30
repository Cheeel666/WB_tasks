package handler

import (
	"dev11/pkg/model"
	"encoding/json"
	"net/http"
)

type result struct {
	Result map[string][]string `json:"result"`
}

// CreateEvent Handler creates event
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))

		return
	}

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	// Вызываем метод создания ивента
	cal := new(model.Calendar)
	err = cal.CreateEvent(event.ID, event.Name, event.Event, event.Dt)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result":"Event created"}`))
}

// UpdateEvent handler updates event
func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	cal := new(model.Calendar)
	err = cal.UpdateEvent(event.ID, event.Event)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result":"Event updated"}`))
}

// DeleteEvent - deletes events
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	cal := new(model.Calendar)
	err = cal.DeleteEvent(event.ID)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result":"Event deleted"}`))
}

// EventsPerDay - returs events  per day
func (h *Handler) EventsPerDay(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}
	cal := new(model.Calendar)
	// TODO parse params
	res, err := cal.FindEvents("day")

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	jsonRes, err := json.Marshal(result{Result: res})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

// EventsPerWeek - returns events per week
func (h *Handler) EventsPerWeek(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}
	cal := new(model.Calendar)
	// TODO parse params
	res, err := cal.FindEvents("week")

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	m := result{res}
	jsonRes, err := json.Marshal(m)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

// EventsPerMonth - returns events per month
func (h *Handler) EventsPerMonth(w http.ResponseWriter, r *http.Request) {
	// Отсекаем неправильные запросы
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
		return
	}
	cal := new(model.Calendar)
	// TODO parse params
	res, err := cal.FindEvents("month")

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	m := result{res}
	jsonRes, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
