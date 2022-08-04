package rest

import (
	"github.com/Hikarikun92/go-example/util"
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
	dto := &LoginDto{}
	if err := util.ParseJson(r.Body, dto); err != nil {
		http.Error(w, err.Error(), err.Status)
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
