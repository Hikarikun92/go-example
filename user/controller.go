package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Controller interface {
	FindAll(w http.ResponseWriter, request *http.Request)
}

func NewController(service Service) Controller {
	return &controllerImpl{service: service}
}

type controllerImpl struct {
	service Service
}

func (c *controllerImpl) FindAll(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := c.service.FindAll()

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonBytes)
}
