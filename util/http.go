package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpError struct {
	error  string
	Status int
}

func (h *HttpError) Error() string {
	return h.error
}

func NewHttpError(error string, status int) *HttpError {
	return &HttpError{error: error, Status: status}
}

func ParseJson(reader io.Reader, dto interface{}) *HttpError {
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return NewHttpError(err.Error(), http.StatusInternalServerError)
	}

	if err = json.Unmarshal(bodyBytes, dto); err != nil {
		return NewHttpError(err.Error(), http.StatusBadRequest)
	}

	return nil
}
