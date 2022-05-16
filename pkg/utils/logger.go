package utils

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	ctxRequestID uint64 = iota
)

type Logger struct {
	logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{*logrus.New()}
}

func SetRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := ctxRequestID
		w.Header().Set("X-Request-ID", strconv.FormatUint(id, 10))
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), id, id)))
	})
}

func (lg *Logger) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := lg.WithFields(logrus.Fields{
			"remote_addr":    r.RemoteAddr,
			"remote_url":     r.RequestURI,
			"request_method": r.Method,
			"request_id":     r.Context().Value(ctxRequestID),
		})
		logger.Infof("start request: id:%s, Method:%s, url:%s", w.Header().Get("X-Request-ID"), r.Method, r.RequestURI)

		resWriter := &responseWriter{w, http.StatusOK}
		start := time.Now()
		next.ServeHTTP(resWriter, r)

		var level logrus.Level
		switch {
		case resWriter.code >= 500:
			level = logrus.ErrorLevel
		case resWriter.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed request: id:%s, %d %s in %v",
			w.Header().Get("X-Request-ID"),
			resWriter.code,
			http.StatusText(resWriter.code),
			time.Since(start),
		)
	})
}
