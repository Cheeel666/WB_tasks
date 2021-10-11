package handler

import (
	"dev11/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	body, err := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal) // check for errors

	event.UserID, err = strconv.Atoi(keyVal["user_id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}

	event.Dt, err = time.Parse("2006-01-02", keyVal["date"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}

	event.Event = keyVal["event"]
	fmt.Println("event:", event.Event)
	h.cal.CreateEvent(event)
	fmt.Println(h.cal)
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
	body, err := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal) // check for errors

	event.ID, err = strconv.Atoi(keyVal["id"])
	event.Event = keyVal["event"]
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}

	h.cal.UpdateEvent(event)
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
	body, err := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal) // check for errors

	event.ID, err = strconv.Atoi(keyVal["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Bad request(bad data)"}`))
		return
	}

	h.cal.DeleteEvent(event)
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

	// TODO parse params
	res, err := h.cal.FindEvents("day", time.Now())

	if res != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	jsonRes, err := json.Marshal(res)
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
	res, err := h.cal.FindEvents("week", time.Now())

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	jsonRes, err := json.Marshal(res)

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

	res, err := h.cal.FindEvents("month", time.Now())

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 error в случ ошибки BI
		w.Write([]byte(`{"error":"BI error"}`))

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
