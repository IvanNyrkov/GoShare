package snippers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// StatusResponse responds with only status code
func StatusResponse(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// StringResponse responds with status code and plain text
func StringResponse(w http.ResponseWriter, code int, s string) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := w.Write([]byte(s))
	return err
}

// JSONResponse responds with status code and JSON data
func JSONResponse(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// XMLResponse responds with status code and XML data
func XMLResponse(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/xml")
	b, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
