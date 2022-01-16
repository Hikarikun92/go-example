package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_controllerImpl_FindAll_withSuccess(t *testing.T) {
	facade := mockFacade{findAllImpl: func() ([]*UserReadDto, error) {
		return []*UserReadDto{
			{Id: 1, Username: "Administrator"},
			{Id: 2, Username: "John Doe"},
			{Id: 3, Username: "Mary Doe"},
		}, nil
	}}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	controller.FindAll(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusOK)
	}

	jsonBytes, _ := ioutil.ReadAll(response.Body)
	contentLength := int(response.ContentLength)
	if contentLength != len(jsonBytes) {
		t.Errorf("Got %v, wanted %v", contentLength, len(jsonBytes))
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Got %v, wanted application/json", contentType)
	}

	json := string(jsonBytes)
	expected := "[{\"id\":1,\"username\":\"Administrator\"},{\"id\":2,\"username\":\"John Doe\"},{\"id\":3,\"username\":\"Mary Doe\"}]"
	if json != expected {
		t.Errorf("Got %v, wanted %v", json, expected)
	}
}

type mockFacade struct {
	findAllImpl func() ([]*UserReadDto, error)
}

func (f *mockFacade) FindAll() ([]*UserReadDto, error) {
	return f.findAllImpl()
}

func Test_controllerImpl_FindAll_withError(t *testing.T) {
	facade := mockFacade{findAllImpl: func() ([]*UserReadDto, error) { return nil, errors.New("Error finding users") }}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	controller.FindAll(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusInternalServerError)
	}
}
