package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-example/security"
	"go-example/user"
	"go-example/util"
	"io/ioutil"
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
	credentials, ok := request.Context().Value("credentials").(*user.Credentials)
	if !ok {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}
	if !util.ContainsString(credentials.Roles, security.ROLE_USER) {
		http.Error(w, "Unauthorized access", http.StatusForbidden)
		return
	}

	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dto := &CreatePostDto{}
	if err = json.Unmarshal(bodyBytes, dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
