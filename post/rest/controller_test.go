package rest

import (
	"github.com/gorilla/mux"
	commentRest "go-example/comment/rest"
	"go-example/user/rest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_FindByUserId_wrongUserId(t *testing.T) {
	facade := mockFacade{}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "asdf"})

	w := httptest.NewRecorder()
	controller.FindByUserId(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusBadRequest)
	}
}

type mockFacade struct {
	findByUserIdImpl func(userId int) []*PostByUserDto
	findByIdImpl     func(id int) *PostByIdDto
}

func (f *mockFacade) FindByUserId(userId int) []*PostByUserDto {
	return f.findByUserIdImpl(userId)
}

func (f *mockFacade) FindById(id int) *PostByIdDto {
	return f.findByIdImpl(id)
}

func Test_FindByUserId_withSuccess(t *testing.T) {
	dtos := []*PostByUserDto{
		{Id: 1, Title: "Example post no. 1", Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PublishedDate: "2021-01-01T12:03:18"},
		{Id: 2, Title: "Another example post", Body: "Integer malesuada lorem non nunc.", PublishedDate: "2021-03-15T17:53:07"},
	}

	facade := mockFacade{findByUserIdImpl: func(userId int) []*PostByUserDto { return dtos }}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "13"})

	w := httptest.NewRecorder()
	controller.FindByUserId(w, req)

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
	expected := "[" +
		"{\"id\":1,\"title\":\"Example post no. 1\",\"body\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit.\",\"publishedDate\":\"2021-01-01T12:03:18\"}," +
		"{\"id\":2,\"title\":\"Another example post\",\"body\":\"Integer malesuada lorem non nunc.\",\"publishedDate\":\"2021-03-15T17:53:07\"}" +
		"]"
	if json != expected {
		t.Errorf("Got %v, wanted %v", json, expected)
	}
}

func Test_FindById_wrongId(t *testing.T) {
	facade := mockFacade{}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "asdf"})

	w := httptest.NewRecorder()
	controller.FindById(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusBadRequest)
	}
}

func Test_FindById_notFound(t *testing.T) {
	facade := mockFacade{findByIdImpl: func(id int) *PostByIdDto { return nil }}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "18"})

	w := httptest.NewRecorder()
	controller.FindById(w, req)

	response := w.Result()
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Got status %v, wanted status %v", response.StatusCode, http.StatusNotFound)
	}
}

func Test_FindById_success(t *testing.T) {
	facade := mockFacade{findByIdImpl: func(id int) *PostByIdDto {
		return &PostByIdDto{
			Id:            id,
			Title:         "Test post",
			Body:          "Some example post",
			PublishedDate: "2019-11-23T10:42:36",
			User:          &rest.UserReadDto{Id: 4, Username: "Some user"},
			Comments: []*commentRest.CommentReadDto{
				{
					Id:            1,
					Title:         "Example comment",
					Body:          "This is great!",
					PublishedDate: "2022-01-15T22:43:56",
					User:          &rest.UserReadDto{Id: 5, Username: "Another user"},
				},
			},
		}
	}}
	controller := controllerImpl{facade: &facade}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "18"})

	w := httptest.NewRecorder()
	controller.FindById(w, req)

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
	expected := "{\"id\":18,\"title\":\"Test post\",\"body\":\"Some example post\",\"publishedDate\":\"2019-11-23T10:42:36\",\"user\":{\"id\":4,\"username\":\"Some user\"}," +
		"\"comments\":[{\"id\":1,\"title\":\"Example comment\",\"body\":\"This is great!\",\"publishedDate\":\"2022-01-15T22:43:56\",\"user\":{\"id\":5,\"username\":\"Another user\"}}]}"
	if json != expected {
		t.Errorf("Got %v, wanted %v", json, expected)
	}
}
