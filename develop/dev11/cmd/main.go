package main

import (
	"net/http"

	"github.com/gmohmad/wbL2/develop/dev11/internal/http/handlers"
)

func main() {
	ch := handlers.NewCalendarHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", ch.CreateEvent)
	mux.HandleFunc("/update_event", ch.UpdateEvent)
	mux.HandleFunc("/delete_event", ch.DeleteEvent)
	mux.HandleFunc("/events_for_day", ch.EventsForDay)
	mux.HandleFunc("/events_for_week", ch.EventsForWeek)
	mux.HandleFunc("/events_for_month", ch.EventsForMonth)

	loggerMux := handlers.NewLogger(mux)

	http.ListenAndServe("localhost:8000", loggerMux)
}
