package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Controller interface {
	Login(w http.ResponseWriter, request *http.Request)
}

type controllerImpl struct {
	facade Facade
}

func NewController(facade Facade) Controller {
	return &controllerImpl{facade: facade}
}

func (c *controllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dto := &LoginDto{}
	if err = json.Unmarshal(bodyBytes, dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := c.facade.Login(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenBytes := []byte(token)
	w.Header().Add("Content-Length", strconv.Itoa(len(tokenBytes)))
	w.Header().Add("Content-Type", "text/plain")
	_, _ = w.Write(tokenBytes)
}
