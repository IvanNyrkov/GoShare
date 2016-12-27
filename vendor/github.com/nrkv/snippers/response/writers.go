package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// Status responds with only status code
func Status(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

// String responds with status code and plain text
func String(w http.ResponseWriter, code int, s string) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := w.Write([]byte(s))
	return err
}

// JSON responds with status code and JSON data
func JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// XML responds with status code and XML data
func XML(w http.ResponseWriter, code int, data interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/xml")
	b, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// File responds with status code and File data
func File(w http.ResponseWriter, r *http.Request, code int, filePath string) {
	w.WriteHeader(code)
	http.ServeFile(w, r, filePath)
}
