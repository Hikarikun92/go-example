package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/Hikarikun92/go-example/security"
	"github.com/Hikarikun92/go-example/util"
	"net/http"
	"strconv"
)

type Controller interface {
	FindByUserId(w http.ResponseWriter, request *http.Request)
	FindById(w http.ResponseWriter, request *http.Request)
	Create(w http.ResponseWriter, request *http.Request)
}

func NewController(facade Facade) Controller {
	return &controllerImpl{facade: facade}
}

type controllerImpl struct {
	facade Facade
}

func (c *controllerImpl) FindByUserId(w http.ResponseWriter, request *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(request)["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := c.facade.FindByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonBytes)
}

func (c *controllerImpl) FindById(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := c.facade.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.NotFound(w, request)
		return
	}

	jsonBytes, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Length", strconv.Itoa(len(jsonBytes)))
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonBytes)
}

func (c *controllerImpl) Create(w http.ResponseWriter, request *http.Request) {
	credentials, credentialsErr := security.AssertRole(request, security.ROLE_USER)
	if credentialsErr != nil {
		http.Error(w, credentialsErr.Error(), credentialsErr.Status)
		return
	}

	dto := &CreatePostDto{}
	if err := util.ParseJson(request.Body, dto); err != nil {
		http.Error(w, err.Error(), err.Status)
		return
	}

	id, err := c.facade.Create(dto, credentials.User.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", "/posts/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}
