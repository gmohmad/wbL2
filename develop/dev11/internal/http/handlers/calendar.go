package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fossoreslp/go-uuid-v4"
	"github.com/gmohmad/wbL2/develop/dev11/internal/calendar"
	"github.com/gmohmad/wbL2/develop/dev11/internal/http/utils"
)

type CalendarHandler struct {
	calendar *calendar.Calendar
}

type ResponseEvents struct {
	Events []calendar.Event `json:"events"`
}

func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{calendar.NewCalendar()}
}

func (c *CalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("POST", w, r)

	utils.CheckContentType(w, r)

	err := r.ParseForm()
	if err != nil {
		utils.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		utils.Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
		return
	}

	eventData := &calendar.EventData{Name: r.FormValue("name"), Date: date}

	c.calendar.CreateEvent(eventData)

	utils.OKMessage(w, "event successfully created")
}

func (c *CalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("POST", w, r)

	utils.CheckContentType(w, r)

	err := r.ParseForm()
	if err != nil {
		utils.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		utils.Error(w, "invalid data: error parsing id", http.StatusBadRequest)
		return
	}

	var date time.Time

	if r.FormValue("date") != "" {
		date, err = time.Parse("2006-01-02", r.FormValue("date"))
		if err != nil {
			utils.Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
			return
		}
	}

	eventData := &calendar.EventData{Name: r.FormValue("name"), Date: date}

	err = c.calendar.UpdateEvent(id, eventData)
	if err != nil {
		utils.Error(w, err.Error(), http.StatusNotFound)
	}

	utils.OKMessage(w, "event successfully updated")
}

func (c *CalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("POST", w, r)

	utils.CheckContentType(w, r)

	err := r.ParseForm()
	if err != nil {
		utils.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		utils.Error(w, "invalid data: error parsing id", http.StatusBadRequest)
		return
	}

	err = c.calendar.DeleteEvent(id)
	if err != nil {
		utils.Error(w, err.Error(), http.StatusNotFound)
	}

	utils.OKMessage(w, "event successfully deleted")
}

func (c *CalendarHandler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("GET", w, r)

	qDate := r.URL.Query().Get("date")
	var date time.Time
	var err error

	if qDate == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", qDate)
		if err != nil {
			utils.Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
			return
		}
	}

	data, err := json.Marshal(ResponseEvents{c.calendar.EventsForDay(date)})

	if err != nil {
		utils.Error(w, "error marshalling events", http.StatusInternalServerError)
		return
	}

	utils.OKData(w, data)
	log.Println("sent events for day")
}

func (c *CalendarHandler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("GET", w, r)

	qDate := r.URL.Query().Get("date")
	var date time.Time
	var err error

	if qDate == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", qDate)
		if err != nil {
			utils.Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
			return
		}
	}

	data, err := json.Marshal(ResponseEvents{c.calendar.EventsForWeek(date)})

	if err != nil {
		utils.Error(w, "error marshalling events", http.StatusInternalServerError)
		return
	}

	utils.OKData(w, data)
	log.Println("sent events for day")
}

func (c *CalendarHandler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	utils.CheckMehod("GET", w, r)

	qDate := r.URL.Query().Get("date")
	var date time.Time
	var err error

	if qDate == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", qDate)
		if err != nil {
			utils.Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
			return
		}
	}

	data, err := json.Marshal(ResponseEvents{c.calendar.EventsForMonth(date)})

	if err != nil {
		utils.Error(w, "error marshalling events", http.StatusInternalServerError)
		return
	}

	utils.OKData(w, data)
	log.Println("sent events for day")
}
