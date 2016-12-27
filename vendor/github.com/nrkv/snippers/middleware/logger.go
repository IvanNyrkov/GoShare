package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"bytes"
	"os"
	"io"
)

// LoggerFormat is a struct that specifies format of printed log unit
type LoggerFormat struct {
	Format string // classic printf-style format
	Value  string // predefined type or hardcoded value
}

const (
	// LoggerTime will be replaced with time of request creation
	LoggerTime    = "{time}"
	// LoggerMethod will be replaced with request method
	LoggerMethod  = "{method}"
	// LoggerURI will be replaced with request URI
	LoggerURI     = "{uri}"
	// LoggerStatus will be replaced with response status
	LoggerStatus  = "{status}"
	// LoggerLatency will be replaced with duration between request and response
	LoggerLatency = "{latency}"
)

// DefaultLoggerConfig is a config that setups log to following format:
// TIME RFC3339 | METHOD | STATUS | LATENCY | URI
// Example:
// 2016/12/24 17:05:57 | GET | 200 | 3.903567ms | /api/accounts/42
var DefaultLoggerConfig = []LoggerFormat{
	{"%s", LoggerTime},
	{"%s", " | "},
	{"%-6s", LoggerMethod},
	{"%s", " | "},
	{"%s", LoggerStatus},
	{"%s", " | "},
	{"%15s", LoggerLatency},
	{"%s", " | "},
	{"%s", LoggerURI},
	{"%s", "\n"},
}

// Logger is a middleware handler that prints info about request (latency, status, uri, method, time)
func Logger(inner http.Handler, out io.Writer, logFormat []LoggerFormat) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()
		// Serve request
		loggedResponse := &loggedResponse{
			ResponseWriter: w,
			status:         200,
		}
		inner.ServeHTTP(loggedResponse, r)
		// Assemble log record
		var format bytes.Buffer
		var values []interface{}
		for _, f := range logFormat {
			switch f.Value {
			case LoggerTime:
				format.WriteString(f.Format)
				values = append(values, start.Format("2006/01/02 15:04:05"))
			case LoggerMethod:
				format.WriteString(f.Format)
				values = append(values, r.Method)
			case LoggerURI:
				format.WriteString(f.Format)
				values = append(values, r.RequestURI)
			case LoggerStatus:
				format.WriteString(f.Format)
				values = append(values, strconv.Itoa(loggedResponse.status))
			case LoggerLatency:
				format.WriteString(f.Format)
				values = append(values, time.Since(start).String())
			default:
				format.WriteString(f.Format)
				values = append(values, f.Value)
			}
		}
		// Print log record
		if out == nil {
			out = os.Stdout
		}
		fmt.Fprintf(out, format.String(), values...)
	})
}

// loggedResponse is a wrapper over the ResponseWriter interface that allows us to use its response status
type loggedResponse struct {
	http.ResponseWriter
	status int
}

// WriteHeader wraps default ResponseWriter.WriteHeader with saving response status
func (l *loggedResponse) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}
