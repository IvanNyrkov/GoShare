package middleware

import (
	"fmt"
	"net/http"
	"time"
)

var LogsBufferSize = 100 // buffer size of log record channel

// logRecord is a struct that describes a log record to be printed
type logRecord struct {
	Start time.Time
	Method string
	Status int
	URI string
	Latency time.Duration
}

// loggedResponseWriter is a wrapper over the ResponseWriter interface that allows us to use its response status
type loggedResponseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader wraps default ResponseWriter.WriteHeader with saving response status
func (l *loggedResponseWriter) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

// Logger is a middleware handler that prints info about request (latency, status, uri, method, time)
// It creates and listens buffered channel for non-blocking log printing
func Logger(inner http.Handler) http.Handler {
	logsChan := make(chan *logRecord, LogsBufferSize)
	go listenLogs(logsChan)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Start timer
		lrw := &loggedResponseWriter{
			ResponseWriter: w,
			status:         200, // by default status is 200, as it is in http.ResponseWriter
		}
		inner.ServeHTTP(lrw, r) // Serve request with status saving
		logsChan <- &logRecord{
				Start: start,
				Method: r.Method,
				Status: lrw.status,
				Latency: time.Since(start),
				URI: r.RequestURI,
		}
	})
}

// listenLogs obtains log records from the queue and prints them
func listenLogs(logsChan chan *logRecord) {
	for logRec := range logsChan {
		fmt.Printf("%s | %-6s | %d | %15s | %s \n",
			logRec.Start.Format("2006/01/02 15:04:05"),
			logRec.Method,
			logRec.Status,
			logRec.Latency.String(),
			logRec.URI,
		)
	}
}
