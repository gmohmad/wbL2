package calendar

import (
	"time"

	"github.com/fossoreslp/go-uuid-v4"
)

type Event struct {
	Id uuid.UUID `json:"id"`
	EventData
}

type EventData struct {
	Name string    `json:"name"`
	Date time.Time `json:"time"`
}

func NewEvent(eventData *EventData) *Event {
	id, _ := uuid.New()
	return &Event{id, *eventData}
}
