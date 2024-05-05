package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LogPayload struct {
	Name  string
	Data  any
	Field string
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}
	errMessage := strings.Split(err.Error(), "desc = ")
	var payload jsonResponse
	payload.Error = true
	payload.Message = errMessage[1]
	return WriteJSON(w, statusCode, payload)
}

func LogInformation(w http.ResponseWriter, data LogPayload) {
	// create some json we'll send in the log microservice
	jsonData, _ := json.MarshalIndent(data, "", "\t")

	// call the service

	request, err := http.NewRequest("POST", "http://logger-service:8001/logs", bytes.NewBuffer(jsonData))
	if err != nil {
		ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()
}
