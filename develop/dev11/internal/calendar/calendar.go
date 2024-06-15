package calendar

import (
	"fmt"
	"sync"
	"time"

	"github.com/fossoreslp/go-uuid-v4"
)

type Calendar struct {
	mu     sync.RWMutex
	events map[uuid.UUID]EventData
}

func NewCalendar() *Calendar {
	return &Calendar{sync.RWMutex{}, make(map[uuid.UUID]EventData)}
}

func (c *Calendar) getEvent(id uuid.UUID) (*EventData, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	event, ok := c.events[id]
	if !ok {
		return nil, fmt.Errorf("event with id '%v' does not exist", id)
	}
	return &event, nil

}

func (c *Calendar) CreateEvent(eventData *EventData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEvent := NewEvent(eventData)
	c.events[newEvent.Id] = newEvent.EventData
}

func (c *Calendar) UpdateEvent(id uuid.UUID, eventData *EventData) error {
	event, err := c.getEvent(id)
	if err != nil {
		return err
	}

	if !eventData.Date.IsZero() {
		event.Date = eventData.Date
	}

	if eventData.Name != "" {
		event.Name = eventData.Name
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.events[id] = *event
	return nil
}

func (c *Calendar) DeleteEvent(id uuid.UUID) error {
	_, err := c.getEvent(id)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.events, id)
	return nil
}

func (c *Calendar) EventsForDay(date time.Time) []Event {
	var events []Event

	cyear, cmonth, cday := date.Date() // current year, month and day

	c.mu.RLock()
	for id, eventData := range c.events {
		year, month, day := eventData.Date.Date()
		if cyear == year && cmonth == month && cday == day {
			events = append(events, Event{id, eventData})
		}
	}
	c.mu.RUnlock()

	return events
}
func (c *Calendar) EventsForWeek(date time.Time) []Event {
	var events []Event

	cyear, cweek := date.ISOWeek() // current year and week

	c.mu.RLock()
	for id, eventData := range c.events {
		year, week := eventData.Date.ISOWeek()
		if cyear == year && cweek == week {
			events = append(events, Event{id, eventData})
		}
	}
	c.mu.RUnlock()

	return events
}
func (c *Calendar) EventsForMonth(date time.Time) []Event {
	var events []Event

	cyear, cmonth, _ := date.Date() // current year and month

	c.mu.RLock()
	for id, eventData := range c.events {
		year, month, _ := eventData.Date.Date()
		if cyear == year && cmonth == month {
			events = append(events, Event{id, eventData})
		}
	}
	c.mu.RUnlock()

	return events
}
