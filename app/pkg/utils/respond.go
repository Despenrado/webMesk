package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

// Respond write respond to network channel
func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// WriteHeader add statuscode to response
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Error response error
func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Printf("ERROR: request_id:%s,%v\n", w.Header().Get("X-Request-ID"), err)
	Respond(w, r, code, map[string]string{"error": err.Error()})
}
