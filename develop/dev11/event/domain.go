package event

import (
	"time"
)

type User struct {
	ID     uint64  `json:"id,omitempty"`
	Events []Event `json:"events,omitempty"`
}

type Event struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateEventRequest struct {
	UserID    uint64    `json:"user_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type UpdateEventRequest struct {
	EventID   uint64    `json:"event_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
