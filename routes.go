package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func loggingRoute(r *http.Request) {
	log.Output(1, fmt.Sprintf("%s - %s - %s - %s - Headers: %+v", r.Method, r.URL, r.Proto, r.RemoteAddr, r.Header))
}

func routeFunc(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggingRoute(r)
		fn(w, r)
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	content, err := json.Marshal(map[string]interface{}{"status": 200, "text": "Hello World"})
	errPanic(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
