package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_controllerImpl_Login_withSuccess(t *testing.T) {
	facade := &mockFacade{loginImpl: func(dto *LoginDto) (string, error) { return "example-token", nil }}
	controller := controllerImpl{facade: facade}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"username\":\"Some user\",\"password\":\"mypass\"}"))
	w := httptest.NewRecorder()
	controller.Login(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusOK)
	}

	tokenBytes, _ := ioutil.ReadAll(response.Body)
	contentLength := int(response.ContentLength)
	if contentLength != len(tokenBytes) {
		t.Errorf("Got %v, wanted %v", contentLength, len(tokenBytes))
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("Got %v, wanted text/plain", contentType)
	}

	token := string(tokenBytes)
	if token != "example-token" {
		t.Errorf("Got %v, wanted example-token", token)
	}
}

type mockFacade struct {
	loginImpl func(dto *LoginDto) (string, error)
}

func (f *mockFacade) Login(dto *LoginDto) (string, error) {
	return f.loginImpl(dto)
}

func Test_controllerImpl_Login_invalidJson(t *testing.T) {
	controller := controllerImpl{}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("some invalid JSON"))
	w := httptest.NewRecorder()
	controller.Login(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusBadRequest)
	}
}

func Test_controllerImpl_Login_loginError(t *testing.T) {
	facade := &mockFacade{loginImpl: func(dto *LoginDto) (string, error) { return "", errors.New("Some error") }}
	controller := controllerImpl{facade: facade}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\"username\":\"Some user\",\"password\":\"mypass\"}"))
	w := httptest.NewRecorder()
	controller.Login(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusUnauthorized)
	}
}
