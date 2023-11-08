package repository

import (
	"calendar/event"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type EventRepository struct {
}

func (eRepo *EventRepository) Create(user_id uint64, e event.Event, storage *event.LocalStorage) error {
	var generatedEventID uint64
	r := rand.New(rand.NewSource(99))
	generatedEventID = r.Uint64()
	for {
		if _, ok := storage.UserEvents[user_id]; ok {
			if _, ok := storage.UserEvents[user_id][generatedEventID]; !ok {
				e.ID = generatedEventID
				storage.UserEvents[user_id][generatedEventID] = e
				break
			} else {
				generatedEventID = r.Uint64()
			}
		} else {
			storage.UserEvents[user_id] = map[uint64]event.Event{generatedEventID: e}
			break
		}
	}
	return nil
}

func (eRepo *EventRepository) Update(user_id uint64, e event.Event, storage *event.LocalStorage) error {
	if _, ok := storage.UserEvents[user_id]; ok {
		if _, ok := storage.UserEvents[user_id][e.ID]; !ok {
			storage.UserEvents[user_id][e.ID] = e
		} else {
			err := errors.New(fmt.Sprintf("event id %d doesn`t exist", e.ID))
			return err
		}
	} else {
		err := errors.New(fmt.Sprintf("user id %d doesn`t exist", user_id))
		return err
	}
	return nil
}

func (eRepo *EventRepository) Delete(user_id uint64, event_id uint64, storage *event.LocalStorage) error {
	if _, ok := storage.UserEvents[user_id]; ok {
		if _, ok := storage.UserEvents[user_id][event_id]; !ok {
			delete(storage.UserEvents[user_id], event_id)
		} else {
			err := errors.New(fmt.Sprintf("event id %d doesn`t exist", event_id))
			return err
		}
	} else {
		err := errors.New(fmt.Sprintf("user id %d doesn`t exist", user_id))
		return err
	}
	return nil
}

func (eRepo *EventRepository) GetForDay(user_id uint64, day time.Time, storage *event.LocalStorage) ([]event.Event, error) {
	events := make([]event.Event, 0)
	if _, ok := storage.UserEvents[user_id]; !ok {
		err := errors.New(fmt.Sprintf("user id %d doesn`t exist", user_id))
		return nil, err
	}
	for _, val := range storage.UserEvents[user_id] {
		tStart := day.Sub(val.StartTime)
		tEnd := day.Sub(val.EndTime)
		tdiff := val.EndTime.Sub(val.StartTime)
		if ((tStart > 0 && tStart < 24*time.Hour) || (tEnd < 0 && tEnd > -24*time.Hour)) && (tdiff > 0 && tdiff < 24*time.Hour) {
			events = append(events, val)
		}
	}
	return events, nil
}

func (eRepo *EventRepository) GetForWeek(user_id uint64, week time.Time, storage *event.LocalStorage) ([]event.Event, error) {
	events := make([]event.Event, 0)
	if _, ok := storage.UserEvents[user_id]; !ok {
		err := errors.New(fmt.Sprintf("user id %d doesn`t exist", user_id))
		return nil, err
	}
	for _, val := range storage.UserEvents[user_id] {
		tStart := week.Sub(val.StartTime)
		tEnd := week.Sub(val.EndTime)
		tdiff := val.EndTime.Sub(val.StartTime)
		if ((tStart > 0 && tStart < 7*24*time.Hour) || (tEnd < 0 && tEnd > -7*24*time.Hour)) && (tdiff > 0 && tdiff < 7*24*time.Hour) {
			events = append(events, val)
		}
	}
	return events, nil
}

func (eRepo *EventRepository) GetForMonth(user_id uint64, month time.Time, storage *event.LocalStorage) ([]event.Event, error) {
	events := make([]event.Event, 0)
	if _, ok := storage.UserEvents[user_id]; !ok {
		err := errors.New(fmt.Sprintf("user id %d doesn`t exist", user_id))
		return nil, err
	}
	for _, val := range storage.UserEvents[user_id] {
		tStart := month.Sub(val.StartTime)
		tEnd := month.Sub(val.EndTime)
		tdiff := val.EndTime.Sub(val.StartTime)
		if ((tStart > 0 && tStart < 30*24*time.Hour) || (tEnd < 0 && tEnd > -30*24*time.Hour)) && (tdiff > 0 && tdiff < 30*24*time.Hour) {
			events = append(events, val)
		}
	}
	return events, nil
}
