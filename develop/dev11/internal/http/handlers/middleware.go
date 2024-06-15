package handlers

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	next http.Handler
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: Client - %s | Method - %s | Address - %s", r.RemoteAddr, r.Method, r.URL.Path)

	start := time.Now()

	l.next.ServeHTTP(w, r)

	log.Printf("Request processed. Took %s.", time.Since(start))
}
