package util

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"testing/iotest"
)

type testDto struct {
	Key1 int    `json:"key1"`
	Key2 string `json:"key2"`
}

func Test_ParseJson_errorReading(t *testing.T) {
	dto := &testDto{}
	err := ParseJson(iotest.ErrReader(errors.New("Some error")), dto)
	if err == nil {
		t.Error("Expected an error, got none")
	}

	if err.Error() != "Some error" {
		t.Errorf("Expected Some error, got %v", err.Error())
	}
	if err.Status != http.StatusInternalServerError {
		t.Errorf("Expected %v, got %v", http.StatusInternalServerError, err.Status)
	}
}

func Test_ParseJson_wrongJson(t *testing.T) {
	dto := &testDto{}
	err := ParseJson(strings.NewReader("{wrong json"), dto)
	if err == nil {
		t.Error("Expected an error, got none")
	}

	if err.Status != http.StatusBadRequest {
		t.Errorf("Expected %v, got %v", http.StatusBadRequest, err.Status)
	}
}

func Test_ParseJson_withSuccess(t *testing.T) {
	dto := &testDto{}
	err := ParseJson(strings.NewReader("{\"key1\":12,\"key2\":\"Some value\"}"), dto)
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}

	if dto.Key1 != 12 {
		t.Errorf("Expected %v, got %v", 12, dto.Key1)
	}
	if dto.Key2 != "Some value" {
		t.Errorf("Expected Some value, got %v", dto.Key2)
	}
}
