package event

// UserEvents {key user_id : {value : {key = Event.ID, value Event}}}
type LocalStorage struct {
	UserEvents map[uint64]map[uint64]Event
}
