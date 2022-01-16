package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Controller interface {
	FindAll(w http.ResponseWriter, request *http.Request)
}

func NewController(facade Facade) Controller {
	return &controllerImpl{facade: facade}
}

type controllerImpl struct {
	facade Facade
}

func (c *controllerImpl) FindAll(w http.ResponseWriter, req *http.Request) {
	users, err := c.facade.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonBytes)
}
