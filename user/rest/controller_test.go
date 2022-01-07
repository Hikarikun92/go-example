package rest

import (
	. "go-example/user"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestFindAllWithSuccess(t *testing.T) {
	service := serviceImpl{findAllImpl: func() []*User {
		return []*User{
			{Id: 1, Username: "Administrator"},
			{Id: 2, Username: "John Doe"},
			{Id: 3, Username: "Mary Doe"},
		}
	}}
	controller := controllerImpl{service: &service}

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

type serviceImpl struct {
	findAllImpl func() []*User
}

func (s *serviceImpl) FindAll() []*User {
	return s.findAllImpl()
}
