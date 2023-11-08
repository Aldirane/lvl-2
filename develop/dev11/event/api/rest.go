package api

import (
	"calendar/event"
	"calendar/event/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	eventStore   repository.EventRepository
	localStorage *event.LocalStorage
}

func NewAPI(repository repository.EventRepository, localStorage *event.LocalStorage) API {
	return API{
		eventStore:   repository,
		localStorage: localStorage,
	}
}

func (a *API) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", loggingMiddleware(a.Create))
	mux.HandleFunc("/update_event", loggingMiddleware(a.Update))
	mux.HandleFunc("/delete_event", loggingMiddleware(a.Delete))
	mux.HandleFunc("/events_for_day", loggingMiddleware(a.Get))
	mux.HandleFunc("/events_for_week", loggingMiddleware(a.Get))
	mux.HandleFunc("/events_for_month", loggingMiddleware(a.Get))

	return mux
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("bad method: %s,  method should be post", r.Method))
		return
	}

	createEvent, err := parseCreateEventRequest(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	e := event.Event{
		ID:        1,
		Title:     createEvent.Title,
		StartTime: createEvent.StartTime,
		EndTime:   createEvent.EndTime,
	}

	err = a.eventStore.Create(uint64(createEvent.UserID), e, a.localStorage)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "can't create event")
		return
	}

	sendSuccessResponse(w, "Event created successfully")
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("bad method: %s,  method should be post", r.Method))
		return
	}

	updateEvent, err := parseUpdateEventRequest(r)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	e := event.Event{
		ID:        updateEvent.EventID,
		Title:     updateEvent.Title,
		StartTime: updateEvent.StartTime,
		EndTime:   updateEvent.EndTime,
	}

	err = a.eventStore.Update(uint64(updateEvent.EventID), e, a.localStorage)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "can't update event")
		return
	}

	sendSuccessResponse(w, "Event updated successfully")
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("bad method: %s,  method should be post", r.Method))
		return
	}

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprint("can't parse form"))
		return
	}

	uid := r.FormValue("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil || user_id < 0 {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprint("can't parse user_id or user_id < 0"))
		return
	}

	eid := r.FormValue("id")
	event_id, err := strconv.Atoi(eid)
	if err != nil || event_id < 0 {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprint("can't parse id or id < 0"))
		return
	}

	err = a.eventStore.Delete(uint64(user_id), uint64(event_id), a.localStorage)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "can't delete event")
		return
	}

	sendSuccessResponse(w, "Event deleted successfully")
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("bad method: %s, method should be get", r.Method))
		return
	}

	uid := r.URL.Query().Get("user_id")
	user_id, err := strconv.Atoi(uid)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, fmt.Sprint("can't parse user_id"))
		return
	}

	date := r.URL.Query().Get("date")
	t, err := time.Parse(time.DateTime, date)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "can't parse date, use 2006-01-02 15:04:05 format")
		return
	}

	events := make([]event.Event, 0)
	switch r.URL.Path {
	case "/events_for_day":
		events, err = a.eventStore.GetForDay(uint64(user_id), t, a.localStorage)
	case "/events_for_week":
		events, err = a.eventStore.GetForWeek(uint64(user_id), t, a.localStorage)
	case "/events_for_month":
		events, err = a.eventStore.GetForMonth(uint64(user_id), t, a.localStorage)
	}

	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "can't get events")
		return
	}

	if len(events) == 0 {
		sendErrorResponse(w, http.StatusNotFound, "events not found")
		return
	}
	sendSuccessResponse(w, events)
}

func parseCreateEventRequest(r *http.Request) (*event.CreateEventRequest, error) {
	userID := r.FormValue("user_id")
	title := r.FormValue("title")
	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")

	if userID == "" || title == "" || startTime == "" || endTime == "" {
		return nil, errors.New("error empty field")
	}
	id, err := strconv.Atoi(userID)
	if err != nil || id < 0 {
		return nil, errors.New("cant parse userID or userID bellow 0")
	}
	tStart, err := time.Parse(time.DateTime, startTime)
	if err != nil {
		err = errors.New("can't parse date, use 2006-01-02 15:04:05 format")
		return nil, err
	}
	tfinish, err := time.Parse(time.DateTime, startTime)
	if err != nil {
		err = errors.New("can't parse date, use 2006-01-02 15:04:05 format")
		return nil, err
	}
	tdiff := tfinish.Sub(tStart)
	if tdiff < 0 {
		err = errors.New("end date can`t be sooner than start date")
		return nil, err
	}
	return &event.CreateEventRequest{
		UserID:    uint64(id),
		Title:     title,
		StartTime: tStart,
		EndTime:   tfinish,
	}, nil
}

func parseUpdateEventRequest(r *http.Request) (*event.UpdateEventRequest, error) {
	eventID := r.FormValue("event_id")
	title := r.FormValue("title")
	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")

	if eventID == "" || title == "" || startTime == "" || endTime == "" {
		return nil, errors.New("error empty field")
	}
	id, err := strconv.Atoi(eventID)
	if err != nil || id < 0 {
		return nil, errors.New("cant parse eventID or ID bellow 0")
	}
	tStart, err := time.Parse(time.DateTime, startTime)
	if err != nil {
		err = errors.New("can't parse date, use 2006-01-02 15:04:05 format")
		return nil, err
	}
	tfinish, err := time.Parse(time.DateTime, startTime)
	if err != nil {
		err = errors.New("can't parse date, use 2006-01-02 15:04:05 format")
		return nil, err
	}
	tdiff := tfinish.Sub(tStart)
	if tdiff < 0 {
		err = errors.New("end date can`t be sooner than start date")
		return nil, err
	}
	return &event.UpdateEventRequest{
		EventID:   uint64(id),
		Title:     title,
		StartTime: tStart,
		EndTime:   tfinish,
	}, nil
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	response := event.ErrorResponse{
		Error: message,
	}
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter, result interface{}) {
	response := map[string]interface{}{
		"result": result,
	}
	json.NewEncoder(w).Encode(response)
}
