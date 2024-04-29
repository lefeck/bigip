package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type RequestError struct {
	Code     int      `json:"code,omitempty"`
	Message  string   `json:"message,omitempty"`
	ErrStack []string `json:"errorStack,omitempty"`
}

// NewRequestError unmarshal a RequestError from a HTTP response body.
func NewRequestError(body io.Reader) (*RequestError, error) {
	var reqErr RequestError
	dec := json.NewDecoder(body)
	if err := dec.Decode(&reqErr); err != nil {
		return nil, fmt.Errorf("failed to decode request error: %v", err)
	}
	return &reqErr, nil
}

// Error implements the errors.Error interface
func (err RequestError) Error() string {
	return fmt.Sprintf("%s (code: %d)", err.Message, err.Code)
}

func (err RequestError) String() string {
	buf := bytes.NewBufferString(err.Error())
	for _, es := range err.ErrStack {
		buf.WriteString("\n   " + es)
	}
	return buf.String()
}
