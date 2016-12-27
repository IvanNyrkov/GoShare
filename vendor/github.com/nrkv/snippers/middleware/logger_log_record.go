package middleware

import (
	"bytes"
	"time"
)

// logRecord is a struct that describes a log record to be printed
type logRecord struct {
	Format string
	Values []interface{}
}

// newLogRecord creates new log record by specified values and format
func newLogRecord(start time.Time, method string, status int, uri string, logFormat []LoggerFormat) *logRecord {
	var format bytes.Buffer
	var values []interface{}
	for _, f := range logFormat {
		switch f.Value {
		case LoggerTime:
			format.WriteString(f.Format)
			values = append(values, start.Format("2006/01/02 15:04:05"))
		case LoggerMethod:
			format.WriteString(f.Format)
			values = append(values, method)
		case LoggerURI:
			format.WriteString(f.Format)
			values = append(values, uri)
		case LoggerStatus:
			format.WriteString(f.Format)
			values = append(values, status)
		case LoggerLatency:
			format.WriteString(f.Format)
			values = append(values, time.Since(start).String())
		default:
			format.WriteString(f.Format)
			values = append(values, f.Value)
		}
	}
	return &logRecord{
		Format: format.String(),
		Values: values,
	}
}
