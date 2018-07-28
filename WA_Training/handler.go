package main

import (
	"fmt"
	"net/http"
)

type message struct {
	referer   string
	eventType string
	userAgent string
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	sessionData := extractSessionInfo(r)
	// sessionData.print()
	writeMessage(sessionData)
}

func extractSessionInfo(r *http.Request) message {
	requestUserAgent := r.Header.Get("User-Agent")
	userInput := message{
		referer:   "http://www.google.com/not-real",
		eventType: "Load",
		userAgent: requestUserAgent,
	}
	return userInput
}

func (m message) print() {
	fmt.Printf("user event created - %+v \n", m)
}

func (m message) toString() string {
	msg := `{ "referer":"` + m.referer + `","eventType":"` + m.eventType + `","userAgent":"` + m.userAgent + `" }`
	return msg
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	message := "OK!"
	w.Write([]byte(message))
}
