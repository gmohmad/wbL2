package utils

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

const contentType = "application/x-www-form-urlencoded"

func Error(w http.ResponseWriter, msg string, code int) {
	log.Println(msg)
	http.Error(w, msg, code)
}

func OKMessage(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
	log.Println(msg)
}

func OKData(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CheckMehod(excpected string, w http.ResponseWriter, r *http.Request) {
	if r.Method != excpected {
		log.Println(fmt.Sprintf("method '%s' not allowed", r.Method))
		http.Error(w, fmt.Sprintf("method '%s' not allowed", r.Method), http.StatusMethodNotAllowed)
	}
}

func CheckContentType(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != contentType {
		log.Println("invalid Content-Type")
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
	}
}

func ValidateDate(w http.ResponseWriter, qDate string) (*time.Time, error) {
	var date time.Time
	var err error

	if qDate == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", qDate)
		if err != nil {
			Error(w, "invalid data: error parsing the date", http.StatusBadRequest)
			return nil, err
		}
	}
	return &date, nil
}
