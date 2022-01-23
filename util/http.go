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

func ParseJson(reader io.Reader, dto interface{}) *HttpError {
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return &HttpError{error: err.Error(), Status: http.StatusInternalServerError}
	}

	if err = json.Unmarshal(bodyBytes, dto); err != nil {
		return &HttpError{error: err.Error(), Status: http.StatusBadRequest}
	}

	return nil
}
