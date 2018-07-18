package ezresp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Message to pass
type Message struct {
	Message string `json:"message"`
}

// ResponseMessage for easy handling server message
func ResponseMessage(w http.ResponseWriter, status int, message string) (err error) {
	w.WriteHeader(status)
	log.Printf("[HTTP %d] %s", status, message)
	return json.NewEncoder(w).Encode(&Message{Message: message})
}

// DecodeMessage from input stream
func DecodeMessage(r io.Reader) (message string, err error) {

	m := Message{}
	if err := json.NewDecoder(r).Decode(&m); err != nil {
		return message, err
	}

	return m.Message, nil
}
