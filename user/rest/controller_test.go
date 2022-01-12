package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_controllerImpl_FindAll_withSuccess(t *testing.T) {
	facade := mockFacade{findAllImpl: func() []*UserReadDto {
		return []*UserReadDto{
			{Id: 1, Username: "Administrator"},
			{Id: 2, Username: "John Doe"},
			{Id: 3, Username: "Mary Doe"},
		}
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
	contentLength, _ := strconv.Atoi(response.Header.Get("Content-Length"))
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
	findAllImpl func() []*UserReadDto
}

func (f *mockFacade) FindAll() []*UserReadDto {
	return f.findAllImpl()
}
