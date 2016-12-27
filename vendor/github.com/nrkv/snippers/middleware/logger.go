package middleware

import (
	"fmt"
	"net/http"
	"time"

	"io"
	"os"
)

const logBuffer = 100 // buffer size of log record channel

// Logger is a middleware handler that prints info about request (latency, status, uri, method, time)
func Logger(inner http.Handler, out io.Writer, logFormat []LoggerFormat) http.Handler {
	// Use StdOut for output by default
	if out == nil {
		out = os.Stdout
	}

	// Create and listen buffered channel for non-blocking log printing
	logsChan := make(chan *logRecord, logBuffer)
	go listenLogs(out, logsChan)

	// Handle request and create log record with information about it
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()
		// Serve request with status saving
		lrw := &loggedResponseWriter{
			ResponseWriter: w,
			status:         200,
		}
		inner.ServeHTTP(lrw, r)
		// Create log record by specified format and add it to a printing queue
		logsChan <- newLogRecord(start, r.Method, lrw.status, r.RequestURI, logFormat)
	})
}

// listenLogs obtains log records from the queue and prints them
func listenLogs(out io.Writer, logsChan chan *logRecord) {
	for logRec := range logsChan {
		fmt.Fprintf(out, logRec.Format, logRec.Values...)
	}
}
